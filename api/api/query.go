//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/api"
)

func (s *Server) GetAPIs(ctx context.Context, in *npool.GetAPIsRequest) (*npool.GetAPIsResponse, error) {
	handler, err := api1.NewHandler(ctx,
		api1.WithConds(in.GetConds()),
		api1.WithOffset(in.GetOffset()),
		api1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAPIs",
			"In", in,
			"Error", err,
		)
		return &npool.GetAPIsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAPIs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAPIs",
			"In", in,
			"Error", err,
		)
		return &npool.GetAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAPIsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetDomains(ctx context.Context, in *npool.GetDomainsRequest) (*npool.GetDomainsResponse, error) {
	handler, err := api1.NewHandler(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDomains",
			"In", in,
			"Error", err,
		)
		return &npool.GetDomainsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos, err := handler.GetDomains(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDomains",
			"In", in,
			"Error", err,
		)
		return &npool.GetDomainsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDomainsResponse{
		Infos: infos,
	}, nil
}
