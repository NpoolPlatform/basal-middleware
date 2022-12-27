package api

import (
	"github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	api.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
