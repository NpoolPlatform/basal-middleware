package api

import (
	"context"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	mgrpb "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func DeleteAPI(ctx context.Context, id string) (*mgrpb.API, error) {
	info, err := apicli.DeleteAPI(ctx, id)
	if err != nil {
		return nil, err
	}

	return info, nil
}