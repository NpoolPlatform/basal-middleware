package api

import (
	"context"

	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	"github.com/NpoolPlatform/basal-middleware/pkg/db"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func (h *Handler) UpdateAPI(ctx context.Context) (info *npool.API, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.API.UpdateOneID(*h.ID),
			&crud.Req{
				Protocol:    h.Protocol,
				Method:      h.Method,
				MethodName:  h.MethodName,
				Path:        h.Path,
				PathPrefix:  h.PathPrefix,
				ServiceName: h.ServiceName,
				Exported:    h.Exported,
				Depracated:  h.Deprecated,
				Domains:     h.Domains,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAPI(ctx)
}
