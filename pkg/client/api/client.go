//nolint:dupl
package api

import (
	"context"
	"fmt"
	"time"

	servicename "github.com/NpoolPlatform/basal-middleware/pkg/servicename"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get api connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateAPI(ctx context.Context, in *npool.APIReq) (*npool.API, error) {
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
	return info.(*npool.API), nil
}

func CreateAPIs(ctx context.Context, in []*npool.APIReq) ([]*npool.API, error) {
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
	return infos.([]*npool.API), nil
}

func UpdateAPI(ctx context.Context, in *npool.APIReq) (*npool.API, error) {
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
	return info.(*npool.API), nil
}

func GetAPIs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.API, uint32, error) {
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
	return infos.([]*npool.API), total, nil
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

func GetAPIOnly(ctx context.Context, conds *npool.Conds) (*npool.API, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAPIs(ctx, &npool.GetAPIsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get api only: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get api only: %v", err)
	}
	if len(infos.([]*npool.API)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.API)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.API)[0], nil
}

func ExistAPI(ctx context.Context, id string) (bool, error) {
	_, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistAPI(ctx, &npool.ExistAPIRequest{
			EntID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail exist api: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail exist api: %v", err)
	}
	return true, nil
}

func DeleteAPI(ctx context.Context, id uint32) (*npool.API, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteAPI(ctx, &npool.DeleteAPIRequest{
			Info: &npool.APIReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete api: %v", err)
	}
	return info.(*npool.API), nil
}
