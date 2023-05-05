//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mgrcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	"github.com/google/uuid"
)

func (s *Server) UpdateAPI(ctx context.Context, in *npool.UpdateAPIRequest) (*npool.UpdateAPIResponse, error) {
	var err error

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mgrcli.UpdateAPI(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update api: %v", err.Error())
		return &npool.UpdateAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAPIResponse{
		Info: info,
	}, nil
}
