package api

import (
	"context"
	"encoding/json"
	"fmt"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func Prepare(body string) (interface{}, error) {
	req := []*npool.API{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return req, nil
}

func Apply(ctx context.Context, req interface{}) error {
	apis, ok := req.([]*npool.APIReq)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	if len(apis) == 0 {
		return nil
	}

	serviceName := apis[0].ServiceName
	protocol := apis[0].Protocol
	_key := key(*serviceName, protocol.String())

	err := Lock(_key, protocol.String())
	if err != nil {
		return err
	}
	defer func() {
		_ = Unlock(_key, protocol.String())
	}()

	handler, err := api1.NewHandler(
		ctx,
		api1.WithReqs(apis),
	)
	if err != nil {
		return err
	}
	_, err = handler.CreateAPIs(ctx)
	if err != nil {
		return err
	}

	return nil
}
