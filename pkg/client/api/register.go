package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"

	config "github.com/NpoolPlatform/go-service-framework/pkg/config"
	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	"github.com/go-resty/resty/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	mgrpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	eventpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/event"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"google.golang.org/grpc"
)

func publish(apis []*mgrpb.APIReq) error {
	return pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		req := &eventpb.RegisterAPIsRequest{
			Info: apis,
		}
		return publisher.Update(
			basetypes.MsgID_RegisterAPIsReq.String(),
			nil,
			nil,
			nil,
			req,
		)
	})
}

func muxAPIs(mux *runtime.ServeMux) []*mgrpb.APIReq {
	var apis []*mgrpb.APIReq
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	protocol := mgrpb.Protocol_HTTP

	valueOfMux := reflect.ValueOf(mux).Elem()
	handlers := valueOfMux.FieldByName("handlers")
	methIter := handlers.MapRange()

	for methIter.Next() {
		for i := 0; i < methIter.Value().Len(); i++ {
			pat := methIter.Value().Index(i).FieldByName("pat")
			tmp := reflect.NewAt(pat.Type(), unsafe.Pointer(pat.UnsafeAddr())).Elem()
			str := tmp.MethodByName("String").Call(nil)[0].String()
			method, ok := mgrpb.Method_value[methIter.Key().String()]
			if !ok {
				logger.Sugar().Warnw("muxAPIs", "Method", methIter.Key().String())
				continue
			}
			_method := mgrpb.Method(method)

			apis = append(apis, &mgrpb.APIReq{
				Protocol:    &protocol,
				ServiceName: &serviceName,
				Method:      &_method,
				Path:        &str,
			})
		}
	}

	return apis
}

func grpcAPIs(server grpc.ServiceRegistrar) []*mgrpb.APIReq {
	srvInfo := server.(*grpc.Server).GetServiceInfo()

	var apis []*mgrpb.APIReq
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	protocol := mgrpb.Protocol_GRPC
	method := mgrpb.Method_STREAM

	for key, info := range srvInfo {
		for _, _method := range info.Methods {
			path := fmt.Sprintf("%v/%v", key, _method.Name)
			methodName := _method.Name
			apis = append(apis, &mgrpb.APIReq{
				Protocol:    &protocol,
				ServiceName: &serviceName,
				Method:      &method,
				MethodName:  &methodName,
				Path:        &path,
			})
		}
	}

	return apis
}

func getGatewayRouters(name string) ([]*EntryPoint, error) {
	const leastDomainLen = 2
	domain := strings.SplitN(name, ".", leastDomainLen)
	if len(domain) < leastDomainLen {
		return nil, errors.New("service name must like example.npool.top")
	}

	// provider can use kubernetes or k8s
	url := fmt.Sprintf(
		"http://traefik.kube-system.svc.cluster.local:38080/api/http/routers?provider=kubernetes&search=%v",
		domain[0],
	)

	// internal already set timeout
	resp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	routers := make([]*EntryPoint, 0)
	err = json.Unmarshal(resp.Body(), &routers)
	if err != nil {
		return nil, err
	}

	return routers, nil
}

func Register(mux *runtime.ServeMux) error {
	go reliableRegister(mux)
	return nil
}

func reliableRegister(mux *runtime.ServeMux) {
	apis := muxAPIs(mux)
	var err error
	for {
		<-time.After(5 * time.Second) //nolint
		if err = registerHttp(apis); err == nil {
			break
		}
		logger.Sugar().Warnw("Register", "Error", err)
	}
}

func registerHttp(apis []*mgrpb.APIReq) error { //nolint
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	gatewayRouters, err := getGatewayRouters(serviceName)
	if err != nil {
		return err
	}

	if len(gatewayRouters) == 0 {
		return fmt.Errorf("invalid routers")
	}

	for _, router := range gatewayRouters {
		prefix, err := router.PathPrefix()
		if err != nil {
			return err
		}
		routerPath, err := router.Path()

		logger.Sugar().Info("RouterPath", routerPath)

		if err != nil {
			return err
		}

		exported := true
		for _, _api := range apis {
			logger.Sugar().Info(routerPath, *_api.Path, prefix)
			if !strings.HasPrefix(*_api.Path, "/v1") {
				continue
			}
			if !strings.HasPrefix(*_api.Path, routerPath) {
				continue
			}
			_api.PathPrefix = &prefix
			_api.Exported = &exported
			_api.Domains = append(_api.Domains, router.Domain())
			logger.Sugar().Info("------------", exported, _api.Domains)
		}
	}

	go reliablePublish(apis)
	return nil
}

func reliablePublish(apis []*mgrpb.APIReq) {
	for {
		<-time.After(5 * time.Second) //nolint
		if err := publish(apis); err == nil {
			break
		}
	}
}

func RegisterGRPC(server grpc.ServiceRegistrar) {
	apis := grpcAPIs(server)
	go reliablePublish(apis)
}
