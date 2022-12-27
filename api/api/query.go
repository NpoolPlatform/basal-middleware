//nolint:nolintlint,dupl
package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/const"

	mgrcli "github.com/NpoolPlatform/basal-manager/pkg/client/api"
	mgrpb "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
)

func (s *Server) GetAPIs(ctx context.Context, in *npool.GetAPIsRequest) (*npool.GetAPIsResponse, error) {
	var err error

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	conds := in.GetConds()
	if conds == nil {
		conds = &mgrpb.Conds{}
	}

	infos, total, err := mgrcli.GetAPIs(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorf("fail get apis: %v", err)
		return &npool.GetAPIsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAPIsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
