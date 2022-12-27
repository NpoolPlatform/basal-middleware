//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mgrapi "github.com/NpoolPlatform/basal-manager/api/api"
	mgrcli "github.com/NpoolPlatform/basal-manager/pkg/client/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func (s *Server) CreateAPI(ctx context.Context, in *npool.CreateAPIRequest) (*npool.CreateAPIResponse, error) {
	var err error

	err = mgrapi.Validate(in.GetInfo())
	if err != nil {
		return &npool.CreateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mgrcli.CreateAPI(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create api: %v", err.Error())
		return &npool.CreateAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAPIResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAPIs(ctx context.Context, in *npool.CreateAPIsRequest) (*npool.CreateAPIsResponse, error) {
	var err error

	if len(in.GetInfos()) == 0 {
		return &npool.CreateAPIsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = mgrapi.ValidateMany(in.GetInfos())
	if err != nil {
		return &npool.CreateAPIsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := mgrcli.CreateAPIs(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create apis: %v", err)
		return &npool.CreateAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAPIsResponse{
		Infos: infos,
	}, nil
}
