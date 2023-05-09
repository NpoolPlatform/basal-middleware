//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateAPI(ctx context.Context, in *npool.UpdateAPIRequest) (*npool.UpdateAPIResponse, error) {
	req := in.GetInfo()
	handler, err := api1.NewHandler(ctx,
		api1.WithID(req.ID),
		api1.WithProtocol(req.Protocol),
		api1.WithServiceName(req.ServiceName),
		api1.WithMethod(req.Method),
		api1.WithMethodName(req.MethodName),
		api1.WithPath(req.Path),
		api1.WithPathPrefix(req.PathPrefix),
		api1.WithDomains(&req.Domains),
		api1.WithDeprecated(req.Depracated),
		api1.WithExported(req.Exported),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateAPI(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAPIResponse{
		Info: info,
	}, nil
}
