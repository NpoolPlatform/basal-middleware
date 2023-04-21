package api

import (
	"context"

	"github.com/google/uuid"

	apicli "github.com/NpoolPlatform/basal-manager/pkg/client/api"
	mgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
)

func DeleteAPI(ctx context.Context, id uuid.UUID) (*mgrpb.API, error) {
	info, err := apicli.DeleteAPI(ctx, id.String())
	if err != nil {
		return nil, err
	}

	return info, nil
}
