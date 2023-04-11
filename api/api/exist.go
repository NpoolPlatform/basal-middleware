//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mgrcli "github.com/NpoolPlatform/basal-manager/pkg/client/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	"github.com/google/uuid"
)

func (s *Server) ExistAPI(ctx context.Context, in *npool.ExistAPIRequest) (*npool.ExistAPIResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("ExistAPI", "Error", err)
		return &npool.ExistAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := mgrcli.ExistAPI(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("ExistAPI", "Error", err)
		return &npool.ExistAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAPIResponse{
		Info: exist,
	}, nil
}
