package api

import (
	"context"

	"github.com/NpoolPlatform/basal-middleware/pkg/db"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	entapi "github.com/NpoolPlatform/basal-middleware/pkg/db/ent/api"
)

func GetDomains(ctx context.Context) (domains []string, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		domains, err = cli.
			API.
			Query().
			GroupBy(entapi.FieldServiceName).
			Strings(_ctx)
		return err
	})
	return
}
