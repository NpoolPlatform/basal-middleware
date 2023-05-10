package user

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
	apis, ok := req.([]*npool.API)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	// TODO: here we should run in transaction
	for _, _api := range apis {
		handler, err := api1.NewHandler(
			ctx,
			api1.WithProtocol(&_api.Protocol),
			api1.WithServiceName(&_api.ServiceName),
			api1.WithMethod(&_api.Method),
			api1.WithMethodName(&_api.MethodName),
			api1.WithPath(&_api.Path),
			api1.WithPathPrefix(&_api.PathPrefix),
			api1.WithDomains(&_api.Domains),
			api1.WithDeprecated(&_api.Depracated),
			api1.WithExported(&_api.Exported),
		)
		if err != nil {
			return err
		}
		if _, err := handler.CreateAPI(ctx); err != nil {
			return err
		}
	}

	return nil
}
