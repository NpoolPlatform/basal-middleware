//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/const"

	mgrcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	mgrpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/api"
)

func (s *Server) GetAPIs(ctx context.Context, in *npool.GetAPIsRequest) (*npool.GetAPIsResponse, error) {
	handler, err := api1.NewHandler(ctx,
		api1.WithConds(in.GetConds()),
		api1.WithOffset(in.GetOffset()),
		api1.WithLimit(in.GetLimit()),
	)

	infos, total, err := mgrcli.GetAPIs(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetAPIs", "Error", err)
		return &npool.GetAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAPIsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetDomains(ctx context.Context, in *npool.GetDomainsRequest) (*npool.GetDomainsResponse, error) {
	infos, err := api1.GetDomains(ctx)
	if err != nil {
		logger.Sugar().Errorw("GetDomains", "Error", err)
		return &npool.GetDomainsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetDomainsResponse{
		Infos: infos,
	}, nil
}

func (s *Server) GetAPIOnly(ctx context.Context, in *npool.GetAPIOnlyRequest) (*npool.GetAPIOnlyResponse, error) {
	var err error

	conds := in.GetConds()
	if conds == nil {
		conds = &mgrpb.Conds{}
	}

	info, err := mgrcli.GetAPIOnly(ctx, conds)
	if err != nil {
		logger.Sugar().Errorw("GetAPIOnly", "Error", err)
		return &npool.GetAPIOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAPIOnlyResponse{
		Info: info,
	}, nil
}
