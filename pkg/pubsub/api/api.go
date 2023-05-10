package api

import (
	"context"
	"encoding/json"
	"fmt"

	api1 "github.com/NpoolPlatform/basal-middleware/pkg/mw/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	eventpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/event"
)

func Prepare(body string) (interface{}, error) {
	req := eventpb.RegisterAPIsRequest{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return req.Info, nil
}

func Apply(ctx context.Context, req interface{}) error {
	apis, ok := req.([]*npool.APIReq)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	type APIHandler struct {
		api1.Handler
	}
	handler := &APIHandler{}

	_, err := handler.CreateAPIs(ctx, apis)
	if err != nil {
		return err
	}
	return nil
}
