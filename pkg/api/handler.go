package api

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/basal-middleware/pkg/const"
	crud "github.com/NpoolPlatform/basal-middleware/pkg/crud/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	"github.com/google/uuid"
)

type Handler struct {
	ID          *uuid.UUID
	Protocol    *npool.Protocol
	ServiceName *string
	Method      *npool.Method
	MethodName  *string
	Path        *string
	PathPrefix  *string
	Exported    *bool
	Deprecated  *bool
	Domains     *[]string
	Conds       *crud.Conds
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithProtocol(protocol *npool.Protocol) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if protocol == nil {
			return nil
		}
		switch *protocol {
		case npool.Protocol_HTTP:
		case npool.Protocol_GRPC:
		default:
			return fmt.Errorf("invalid protocol %v: ", protocol)
		}

		h.Protocol = protocol
		return nil
	}
}

func WithServiceName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		const leastNameLen = 2
		if len(*name) < leastNameLen {
			return fmt.Errorf("service name %v too short", *name)
		}

		h.ServiceName = name
		return nil
	}
}

func WithMethod(method *npool.Method) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if method == nil {
			return nil
		}
		switch *method {
		case npool.Method_GET:
		case npool.Method_POST:
		case npool.Method_STREAM:
		default:
			return fmt.Errorf("invalid method %v: ", method)
		}

		h.Method = method
		return nil
	}
}

func WithMethodName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		h.MethodName = name
		return nil
	}
}

func WithPath(path *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if path == nil {
			return nil
		}
		h.Path = path
		return nil
	}
}

func WithPathPrefix(prefix *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if prefix == nil {
			return nil
		}
		h.PathPrefix = prefix
		return nil
	}
}

func WithExported(exported *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if exported == nil {
			return nil
		}
		h.Exported = exported
		return nil
	}
}

func WithDeprecated(deprecated *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deprecated == nil {
			return nil
		}
		h.Deprecated = deprecated
		return nil
	}
}

func WithDomains(domains *[]string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if domains == nil {
			return nil
		}
		h.Domains = domains
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.ID.Op, Val: id}
		}
		if conds.Protocol != nil {
			h.Conds.Protocol = &cruder.Cond{Op: conds.Protocol.Op, Val: conds.Protocol.String()}
		}
		if conds.ServiceName != nil {
			h.Conds.ServiceName = &cruder.Cond{Op: conds.ServiceName.Op, Val: conds.GetServiceName().GetValue()}
		}
		if conds.Method != nil {
			h.Conds.Method = &cruder.Cond{Op: conds.Method.Op, Val: conds.Method.String()}
		}
		if conds.Path != nil {
			h.Conds.Path = &cruder.Cond{Op: conds.Path.Op, Val: conds.Path}
		}
		if conds.Exported != nil {
			h.Conds.Exported = &cruder.Cond{Op: conds.Exported.Op, Val: conds.Exported}
		}
		if conds.Depracated != nil {
			h.Conds.Depracated = &cruder.Cond{Op: conds.Depracated.Op, Val: conds.Depracated}
		}
		if conds.IDs != nil {
			if len(conds.GetIDs().GetValue()) > 0 {
				ids := []uuid.UUID{}
				for _, id := range conds.GetIDs().GetValue() {
					_id, err := uuid.Parse(id)
					if err != nil {
						return err
					}
					ids = append(ids, _id)
				}

				h.Conds.IDs = &cruder.Cond{Op: conds.ID.Op, Val: ids}
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
