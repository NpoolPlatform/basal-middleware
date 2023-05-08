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
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
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
	ret = npool.API{
		Protocol:    npool.Protocol_GRPC,
		ProtocolStr: npool.Protocol_GRPC.String(),
		ServiceName: "basal-middleware.npool.top",
		Method:      npool.Method_POST,
		MethodStr:   npool.Method_POST.String(),
		MethodName:  uuid.NewString(),
		Path:        uuid.NewString(),
		PathPrefix:  "/api/basal-middleware",
		Domains:     []string{"api.npool.top"},
		DomainsStr:  "[\"api.npool.top\"]",
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
		assert.Equal(t, info, &ret)
	}
}

func updateAPI(t *testing.T) {
	ret.Protocol = npool.Protocol_HTTP
	ret.ProtocolStr = npool.Protocol_HTTP.String()
	ret.ServiceName = uuid.NewString()
	ret.Method = npool.Method_STREAM
	ret.MethodStr = npool.Method_STREAM.String()
	ret.MethodName = uuid.NewString()
	ret.Path = uuid.NewString()
	ret.PathPrefix = uuid.NewString()
	ret.Domains = []string{"api.npool.top", "procyon.vip"}
	ret.DomainsStr = "[\"api.npool.top\", \"procyon.vip\"]"
	ret.Exported = true
	ret.Depracated = true

	var (
		req = &npool.APIReq{
			ID:          &ret.ID,
			Protocol:    &ret.Protocol,
			ServiceName: &ret.ServiceName,
			Method:      &ret.Method,
			MethodName:  &ret.MethodName,
			Path:        &ret.Path,
			PathPrefix:  &ret.PathPrefix,
			Domains:     ret.Domains,
			Exported:    &ret.Exported,
			Depracated:  &ret.Depracated,
		}
	)
	info, err := UpdateAPI(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAPIs(t *testing.T) {
	infos, _, err := GetAPIs(context.Background(), &npool.Conds{
		ServiceName: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ServiceName,
		},
		Path: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.Path,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotNil(t, infos)
	}
}

func getAPIOnly(t *testing.T) {
	info, err := GetAPIOnly(context.Background(), &npool.Conds{
		ServiceName: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ServiceName,
		},
		Path: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.Path,
		},
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, info)
	}
}

func existAPI(t *testing.T) {
	exist, _ := ExistAPI(context.Background(), ret.ID)
	assert.True(t, exist)
}

func deleteAPI(t *testing.T) {
	info, err := DeleteAPI(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
	info, err = GetAPIOnly(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	})
	assert.Nil(t, err)
	assert.Nil(t, info)
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
	t.Run("updateAPI", updateAPI)
	t.Run("getAPIs", getAPIs)
	t.Run("getAPIOnly", getAPIOnly)
	t.Run("existAPI", existAPI)
	t.Run("deleteAPI", deleteAPI)
}
