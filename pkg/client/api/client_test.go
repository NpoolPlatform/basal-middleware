package api

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
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
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
	ret = npool.API{
		Protocol:    npool.Protocol_GRPC,
		ServiceName: "basal-middleware.npool.top",
		Method:      npool.Method_POST,
		MethodName:  uuid.NewString(),
		Path:        uuid.NewString(),
		PathPrefix:  "/api/basal-middleware",
		Domains:     []string{"api.npool.top"},
		Exported:    false,
		Depracated:  false,
	}
)

func createAPI(t *testing.T) {
	var (
		req = &npool.APIReq{
			Protocol:    &ret.Protocol,
			ServiceName: &ret.ServiceName,
			Method:      &ret.Method,
			MethodName:  &ret.MethodName,
			Path:        &ret.Path,
			PathPrefix:  &ret.PathPrefix,
			Domains:     ret.Domains,
		}
	)
	info, err := CreateAPI(context.Background(), req)
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
	}
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction { //nolint
		return
	}
	// Here won't pass test due to we always test with localhost
	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAPI", createAPI)
}
