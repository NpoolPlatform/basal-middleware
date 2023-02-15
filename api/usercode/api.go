package usercode

import (
	"github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	usercode.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	usercode.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
