package usercode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	google "github.com/NpoolPlatform/go-service-framework/pkg/google"
	usercode1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/usercode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyUserCode(ctx context.Context, in *npool.VerifyUserCodeRequest) (*npool.VerifyUserCodeResponse, error) {
	handler, err := usercode1.NewHandler(ctx,
		usercode1.WithAppID(&in.AppID),
		usercode1.WithPrefix(&in.Prefix),
		usercode1.WithAccount(&in.Account),
		usercode1.WithCode(&in.Code),
		usercode1.WithAccountType(&in.AccountType),
		usercode1.WithUsedFor(&in.UsedFor),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyUserCode",
			"In", in,
			"Error", err,
		)
		return &npool.VerifyUserCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	switch in.GetAccountType() {
	case basetypes.SignMethod_Email:
		fallthrough //nolint
	case basetypes.SignMethod_Mobile:
		err := handler.VerifyUserCode(ctx)
		if err != nil {
			return &npool.VerifyUserCodeResponse{}, status.Error(codes.Internal, err.Error())
		}
	case basetypes.SignMethod_Google:
		valid, err := google.VerifyCode(in.GetAccount(), in.GetCode())
		if err != nil {
			return &npool.VerifyUserCodeResponse{}, status.Error(codes.Internal, err.Error())
		}
		if !valid {
			return &npool.VerifyUserCodeResponse{}, status.Error(codes.Internal, "GoogleCode is invalid")
		}
	}

	return &npool.VerifyUserCodeResponse{}, nil
}
