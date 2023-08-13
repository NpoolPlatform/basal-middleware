package usercode

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/basal-middleware/pkg/testinit"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.UserCode{
		Prefix:      uuid.NewString(),
		AppID:       uuid.NewString(),
		Account:     uuid.NewString(),
		AccountType: basetypes.SignMethod_Email,
		UsedFor:     basetypes.UsedFor_Signin,
	}
)

func createUserCode(t *testing.T) {
	var (
		req = &npool.CreateUserCodeRequest{
			Prefix:      ret.Prefix,
			AppID:       ret.AppID,
			Account:     ret.Account,
			AccountType: ret.AccountType,
			UsedFor:     ret.UsedFor,
		}
	)
	info, err := CreateUserCode(context.Background(), req)
	if assert.Nil(t, err) {
		ret.Code = info.Code
		ret.NextAt = info.NextAt
		ret.ExpireAt = info.ExpireAt
		assert.Equal(t, info, &ret)
	}
}

func verifyUserCode(t *testing.T) {
	var (
		req = &npool.VerifyUserCodeRequest{
			Prefix:      ret.Prefix,
			AppID:       ret.AppID,
			Account:     ret.Account,
			AccountType: ret.AccountType,
			UsedFor:     ret.UsedFor,
			Code:        ret.Code,
		}
	)
	err := VerifyUserCode(context.Background(), req)
	assert.Nil(t, err)
}
func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	// Here won't pass test due to we always test with localhost
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createUserCode", createUserCode)
	t.Run("verifyUserCode", verifyUserCode)
}
