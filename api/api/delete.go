//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func (s *Server) DeleteAPI(ctx context.Context, in *npool.DeleteAPIRequest) (*npool.DeleteAPIResponse, error) {
	handler, err := api1.NewHandler(ctx,
		api1.WithID(in.GetInfo().ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteAPI(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAPIResponse{
		Info: info,
	}, nil
}
