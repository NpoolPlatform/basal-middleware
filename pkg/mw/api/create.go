package api

import (
	"context"

	"github.com/NpoolPlatform/basal-middleware/pkg/db"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	converter "github.com/NpoolPlatform/basal-middleware/pkg/converter/api"
	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	entapi "github.com/NpoolPlatform/basal-middleware/pkg/db/ent/api"

	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateAPIs(ctx context.Context, in []*npool.APIReq) ([]*npool.API, error) {
	var infos []*npool.API

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, info := range in {
			rets, err := tx.
				API.
				Query().
				Where(
					entapi.Protocol(info.GetProtocol().String()),
					entapi.ServiceName(info.GetServiceName()),
					entapi.Method(info.GetMethod().String()),
					entapi.Path(info.GetPath()),
				).
				All(_ctx)
			if err != nil {
				return err
			}
			if len(rets) > 1 {
				logger.Sugar().Warnw("CreateAPIs", "Rets", rets, "Warn", "> 1")
			}
			if len(rets) > 0 {
				infos = append(infos, converter.Ent2Grpc(rets[0]))
				continue
			}
			info2, err := crud.CreateSet(tx.API.Create(), &crud.Req{
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
			infos = append(infos, converter.Ent2Grpc(info2))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func (h *Handler) CreateAPI(ctx context.Context) (*npool.API, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.API.Create(),
			&crud.Req{
				Protocol:    handler.Protocol,
				Method:      handler.Method,
				MethodName:  handler.MethodName,
				Path:        handler.Path,
				PathPrefix:  handler.PathPrefix,
				ServiceName: handler.ServiceName,
				Domains:     handler.Domains,
				Exported:    handler.Exported,
				Depracated:  handler.Deprecated,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAPI(ctx)
}
