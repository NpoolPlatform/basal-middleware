package api

import (
	"context"

	"github.com/NpoolPlatform/basal-manager/pkg/db"
	"github.com/NpoolPlatform/basal-manager/pkg/db/ent"

	converter "github.com/NpoolPlatform/basal-manager/pkg/converter/api"
	crud "github.com/NpoolPlatform/basal-manager/pkg/crud/api"
	entapi "github.com/NpoolPlatform/basal-manager/pkg/db/ent/api"

	mgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
)

func CreateAPIs(ctx context.Context, in []*mgrpb.APIReq) ([]*mgrpb.API, error) {
	var infos []*mgrpb.API

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
			if len(rets) > 0 {
				infos = append(infos, converter.Ent2Grpc(rets[0]))
				continue
			}

			info2, err := crud.CreateSet(tx.API.Create(), info).Save(_ctx)
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
