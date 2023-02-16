package usercode

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	google "github.com/NpoolPlatform/basal-middleware/pkg/google"
	usercode1 "github.com/NpoolPlatform/basal-middleware/pkg/usercode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) VerifyUserCode(ctx context.Context, in *npool.VerifyUserCodeRequest) (*npool.VerifyUserCodeResponse, error) { //nolint
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.VerifyUserCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetPrefix() == "" {
		return &npool.VerifyUserCodeResponse{}, status.Error(codes.InvalidArgument, "Prefix is invalid")
	}
	if in.GetAccount() == "" {
		return &npool.VerifyUserCodeResponse{}, status.Error(codes.InvalidArgument, "Account is invalid")
	}
	if in.GetCode() == "" {
		return &npool.VerifyUserCodeResponse{}, status.Error(codes.InvalidArgument, "Code is invalid")
	}
	switch in.GetAccountType() {
	case basetypes.SignMethod_Email:
	case basetypes.SignMethod_Mobile:
	case basetypes.SignMethod_Google:
	default:
		return &npool.VerifyUserCodeResponse{}, status.Error(codes.InvalidArgument, "AccountType is invalid")
	}
	switch in.GetUsedFor() {
	case basetypes.UsedFor_Signup:
	case basetypes.UsedFor_Signin:
	case basetypes.UsedFor_Update:
	case basetypes.UsedFor_Contact:
	case basetypes.UsedFor_SetWithdrawAddress:
	case basetypes.UsedFor_Withdraw:
	case basetypes.UsedFor_CreateInvitationCode:
	case basetypes.UsedFor_SetCommission:
	case basetypes.UsedFor_SetTransferTargetUser:
	case basetypes.UsedFor_Transfer:
	default:
		return &npool.VerifyUserCodeResponse{}, status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}

	switch in.GetAccountType() {
	case basetypes.SignMethod_Email:
		fallthrough //nolint
	case basetypes.SignMethod_Mobile:
		err := usercode1.VerifyUserCode(
			ctx,
			in.GetPrefix(),
			in.GetAppID(),
			in.GetAccount(),
			in.GetCode(),
			in.GetAccountType(),
			in.GetUsedFor())
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
