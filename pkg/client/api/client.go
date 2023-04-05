//nolint:dupl
package api

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get api connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateAPI(ctx context.Context, in *mgrpb.APIReq) (*mgrpb.API, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateAPI(ctx, &npool.CreateAPIRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create api: %v", err)
	}
	return info.(*mgrpb.API), nil
}

func CreateAPIs(ctx context.Context, in []*mgrpb.APIReq) ([]*mgrpb.API, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateAPIs(ctx, &npool.CreateAPIsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create apis: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create apis: %v", err)
	}
	return infos.([]*mgrpb.API), nil
}

func UpdateAPI(ctx context.Context, in *mgrpb.APIReq) (*mgrpb.API, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateAPI(ctx, &npool.UpdateAPIRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update api: %v", err)
	}
	return info.(*mgrpb.API), nil
}

func GetAPIs(ctx context.Context, conds *mgrpb.Conds, limit, offset int32) ([]*mgrpb.API, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAPIs(ctx, &npool.GetAPIsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get apis: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get apis: %v", err)
	}
	return infos.([]*mgrpb.API), total, nil
}

func GetDomains(ctx context.Context) ([]string, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDomains(ctx, &npool.GetDomainsRequest{})
		if err != nil {
			return nil, fmt.Errorf("fail get domains: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get domains: %v", err)
	}
	return infos.([]string), nil
}
