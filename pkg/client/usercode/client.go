package usercode

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/usercode"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get usercode connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateUserCode(ctx context.Context, in *npool.CreateUserCodeRequest) (*npool.UserCode, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateUserCode(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail create usercode: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create usercode: %v", err)
	}
	return info.(*npool.UserCode), nil
}

func VerifyUserCode(ctx context.Context, in *npool.VerifyUserCodeRequest) error {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		_, err := cli.VerifyUserCode(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail verify usercode: %v", err)
		}
		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("fail verify usercode: %v", err)
	}
	return nil
}
