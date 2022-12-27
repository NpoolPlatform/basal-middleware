//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/const"

	mgrapi "github.com/NpoolPlatform/basal-manager/api/api"
	mgrcli "github.com/NpoolPlatform/basal-manager/pkg/client/api"
	mgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	"github.com/google/uuid"
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

func (s *Server) GetAPIs(ctx context.Context, in *npool.GetAPIsRequest) (*npool.GetAPIsResponse, error) {
	var err error

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	conds := in.GetConds()
	if conds == nil {
		conds = &mgrpb.Conds{}
	}

	infos, total, err := mgrcli.GetAPIs(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorf("fail get apis: %v", err)
		return &npool.GetAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAPIsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
