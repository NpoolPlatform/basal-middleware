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

func reliableRegister(apis []*mgrpb.APIReq) {
	// Add PubSub Impl
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		req := &eventpb.RegisterAPIsRequest{
			Info: apis,
		}
		logger.Sugar().Info("req", apis)
		return publisher.Update(
			basetypes.MsgID_RegisterAPIsReq.String(),
			nil,
			nil,
			nil,
			req,
		)
	}); err != nil {
		logger.Sugar().Errorw(
			"reliableRegister",
			"APIs", apis,
			"Error", err,
		)
	}
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
	apis := muxAPIs(mux)

	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)

	ticker := time.NewTicker(400 * time.Millisecond)
	done := make(chan bool)
	fmt.Println("Started!")
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Stopped!")

	gatewayRouters, err := getGatewayRouters(serviceName)
	if err != nil {
		return err
	}

	for _, router := range gatewayRouters {
		prefix, err := router.PathPrefix()
		if err != nil {
			return err
		}
		routerPath, err := router.Path()
		if err != nil {
			return err
		}

		exported := true
		for _, _api := range apis {
			if !strings.HasPrefix(*_api.Path, "/v1") {
				continue
			}
			if !strings.HasPrefix(*_api.Path, routerPath) {
				continue
			}
			_api.PathPrefix = &prefix
			_api.Exported = &exported
			_api.Domains = append(_api.Domains, router.Domain())
		}
	}

	go reliableRegister(apis)

	return nil
}

func RegisterGRPC(server grpc.ServiceRegistrar) {
	apis := grpcAPIs(server)
	go reliableRegister(apis)
}
