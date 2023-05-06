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

func (s *Server) ExistAPI(ctx context.Context, in *npool.ExistAPIRequest) (*npool.ExistAPIResponse, error) {
	handler, err := api1.NewHandler(ctx, api1.WithID(&in.ID))
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAPI",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := handler.ExistAPI(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAPI",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAPIResponse{
		Info: exist,
	}, nil
}
