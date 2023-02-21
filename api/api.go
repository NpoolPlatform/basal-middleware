package api

import (
	basal "github.com/NpoolPlatform/message/npool/basal/mw/v1"

	api1 "github.com/NpoolPlatform/basal-middleware/api/api"
	usercode "github.com/NpoolPlatform/basal-middleware/api/usercode"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	basal.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	basal.RegisterMiddlewareServer(server, &Server{})
	api1.Register(server)
	usercode.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
