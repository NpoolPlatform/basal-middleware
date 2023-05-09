package usercode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"

	usercode1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/usercode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserCode(ctx context.Context, in *npool.CreateUserCodeRequest) (*npool.CreateUserCodeResponse, error) {
	handler, err := usercode1.NewHandler(ctx,
		usercode1.WithAppID(&in.AppID),
		usercode1.WithPrefix(&in.Prefix),
		usercode1.WithAccount(&in.Account),
		usercode1.WithAccountType(&in.AccountType),
		usercode1.WithUsedFor(&in.UsedFor),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserCode",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateUserCode(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserCode",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUserCodeResponse{
		Info: info,
	}, nil
}
