//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/api"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (s *Server) CreateAPI(ctx context.Context, in *npool.CreateAPIRequest) (*npool.CreateAPIResponse, error) {
	req := in.GetInfo()
	handler, err := api1.NewHandler(ctx,
		api1.WithProtocol(req.Protocol),
		api1.WithServiceName(req.ServiceName),
		api1.WithMethod(req.Method),
		api1.WithMethodName(req.MethodName),
		api1.WithPath(req.Path),
		api1.WithPathPrefix(req.PathPrefix),
		api1.WithDomains(&req.Domains),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	handler.Conds = &api.Conds{
		Protocol: &cruder.Cond{
			Op:  cruder.EQ,
			Val: req.Protocol.String(),
		},
		ServiceName: &cruder.Cond{
			Op:  cruder.EQ,
			Val: req.ServiceName,
		},
		Method: &cruder.Cond{
			Op:  cruder.EQ,
			Val: req.Method.String(),
		},
		Path: &cruder.Cond{
			Op:  cruder.EQ,
			Val: req.Path,
		},
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

	for _, info := range in.GetInfos() {
		_, err := api1.NewHandler(ctx,
			api1.WithProtocol(info.Protocol),
			api1.WithServiceName(info.ServiceName),
			api1.WithMethod(info.Method),
			api1.WithMethodName(info.MethodName),
			api1.WithPath(info.Path),
			api1.WithPathPrefix(info.PathPrefix),
			api1.WithDomains(&info.Domains),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"CreateAPIs",
				"In", in,
				"Error", err,
			)
			return &npool.CreateAPIsResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	handler, err := api1.NewHandler(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := handler.CreateAPIs(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create apis: %v", err)
		return &npool.CreateAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAPIsResponse{
		Infos: infos,
	}, nil
}
