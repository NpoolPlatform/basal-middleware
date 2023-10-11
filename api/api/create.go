//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/api"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (s *Server) CreateAPI(ctx context.Context, in *npool.CreateAPIRequest) (*npool.CreateAPIResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
		)
		return &npool.CreateAPIResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := api1.NewHandler(
		ctx,
		api1.WithEntID(req.EntID, false),
		api1.WithProtocol(req.Protocol, true),
		api1.WithServiceName(req.ServiceName, true),
		api1.WithMethod(req.Method, true),
		api1.WithMethodName(req.MethodName, true),
		api1.WithPath(req.Path, true),
		api1.WithPathPrefix(req.PathPrefix, true),
		api1.WithDomains(&req.Domains, true),
		api1.WithDeprecated(req.Deprecated, false),
		api1.WithExported(req.Exported, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	handler.Conds = &crud.Conds{
		Protocol:    &cruder.Cond{Op: cruder.EQ, Val: *req.Protocol},
		ServiceName: &cruder.Cond{Op: cruder.EQ, Val: *req.ServiceName},
		Method:      &cruder.Cond{Op: cruder.EQ, Val: *req.Method},
		Path:        &cruder.Cond{Op: cruder.EQ, Val: *req.Path},
	}
	info, err := handler.GetAPIOnly(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAPIResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info != nil {
		return &npool.CreateAPIResponse{
			Info: info,
		}, nil
	}

	info, err = handler.CreateAPI(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAPIResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAPIs(ctx context.Context, in *npool.CreateAPIsRequest) (*npool.CreateAPIsResponse, error) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw(
			"CreateAPIs",
			"In", in,
		)
		return &npool.CreateAPIsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	handler, err := api1.NewHandler(
		ctx,
		api1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := handler.CreateAPIs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAPIsResponse{
		Infos: infos,
	}, nil
}
