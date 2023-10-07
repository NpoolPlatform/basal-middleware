package api

import (
	"context"

	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	"github.com/NpoolPlatform/basal-middleware/pkg/db"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	entapi "github.com/NpoolPlatform/basal-middleware/pkg/db/ent/api"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *Handler) CreateAPIs(ctx context.Context, in []*npool.APIReq) ([]*npool.API, error) {
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, info := range in {
			_api, err := tx.
				API.
				Query().
				Where(
					entapi.Protocol(info.GetProtocol().String()),
					entapi.ServiceName(info.GetServiceName()),
					entapi.Method(info.GetMethod().String()),
					entapi.Path(info.GetPath()),
				).
				Only(_ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return err
				}
			}
			if _api != nil {
				ids = append(ids, _api.EntID)
				continue
			}
			_api, err = crud.CreateSet(tx.API.Create(), &crud.Req{
				Protocol:    info.Protocol,
				Method:      info.Method,
				MethodName:  info.MethodName,
				Path:        info.Path,
				PathPrefix:  info.PathPrefix,
				ServiceName: info.ServiceName,
				Domains:     &info.Domains,
				Depracated:  info.Depracated,
				Exported:    info.Exported,
			}).Save(_ctx)
			if err != nil {
				return err
			}
			ids = append(ids, _api.EntID)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &crud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))
	infos, _, err := h.GetAPIs(ctx)
	return infos, err
}

func (h *Handler) CreateAPI(ctx context.Context) (*npool.API, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.API.Create(),
			&crud.Req{
				Protocol:    h.Protocol,
				Method:      h.Method,
				MethodName:  h.MethodName,
				Path:        h.Path,
				PathPrefix:  h.PathPrefix,
				ServiceName: h.ServiceName,
				Domains:     h.Domains,
				Exported:    h.Exported,
				Depracated:  h.Deprecated,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		h.EntID = &info.EntID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAPI(ctx)
}
