package usercode

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	usercode1 "github.com/NpoolPlatform/basal-middleware/pkg/usercode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) CreateUserCode(ctx context.Context, in *npool.CreateUserCodeRequest) (*npool.CreateUserCodeResponse, error) { //nolint
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return &npool.CreateUserCodeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetPrefix() == "" {
		return &npool.CreateUserCodeResponse{}, status.Error(codes.InvalidArgument, "Prefix is invalid")
	}
	if in.GetAccount() == "" {
		return &npool.CreateUserCodeResponse{}, status.Error(codes.InvalidArgument, "Account is invalid")
	}
	switch in.GetAccountType() {
	case basetypes.SignMethod_Email:
	case basetypes.SignMethod_Mobile:
	default:
		return &npool.CreateUserCodeResponse{}, status.Error(codes.InvalidArgument, "AccountType is invalid")
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
		return &npool.CreateUserCodeResponse{}, status.Error(codes.InvalidArgument, "UsedFor is invalid")
	}

	info, err := usercode1.CreateUserCode(
		ctx,
		in.GetPrefix(),
		in.GetAppID(),
		in.GetAccount(),
		in.GetAccountType(),
		in.GetUsedFor())
	if err != nil {
		return &npool.CreateUserCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUserCodeResponse{
		Info: info,
	}, nil
}
