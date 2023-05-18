package api

import (
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func Ent2Grpc(row *ent.API) *npool.API {
	if row == nil {
		return nil
	}

	return &npool.API{
		ID:          row.ID.String(),
		Protocol:    npool.Protocol(npool.Protocol_value[row.Protocol]),
		ServiceName: row.ServiceName,
		Method:      npool.Method(npool.Method_value[row.Method]),
		Path:        row.Path,
		Exported:    row.Exported,
		PathPrefix:  row.PathPrefix,
		Domains:     row.Domains,
		Depracated:  row.Depracated,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.API) []*npool.API {
	infos := []*npool.API{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
