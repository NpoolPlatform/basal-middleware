//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	"github.com/google/uuid"
)

func (s *Server) DeleteAPI(ctx context.Context, in *npool.DeleteAPIRequest) (*npool.DeleteAPIResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteAPI", "Error", err)
		return &npool.DeleteAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := api1.DeleteAPI(ctx, in.ID)
	if err != nil {
		logger.Sugar().Errorw("DeleteAPI", "Error", err)
		return &npool.DeleteAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAPIResponse{
		Info: info,
	}, nil
}
