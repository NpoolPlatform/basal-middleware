//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateAPI(ctx context.Context, in *npool.UpdateAPIRequest) (*npool.UpdateAPIResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
		)
		return &npool.UpdateAPIResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := api1.NewHandler(ctx,
		api1.WithID(req.ID, true),
		api1.WithDeprecated(req.Deprecated, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAPIResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateAPI(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAPI",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAPIResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAPIResponse{
		Info: info,
	}, nil
}
