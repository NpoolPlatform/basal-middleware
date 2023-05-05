package api

import (
	"fmt"

	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/basal-middleware/pkg/db/ent/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uuid.UUID
	Protocol    *npool.Protocol
	ServiceName *string
	Method      *npool.Method
	MethodName  *string
	Path        *string
	Exported    *bool
	PathPrefix  *string
	Domains     *[]string
	Depracated  *bool
	DeletedAt   *uint32
}

func CreateSet(c *ent.APICreate, req *Req) *ent.APICreate {
	if req.Protocol != nil {
		c.SetProtocol(req.Protocol.String())
	}
	if req.ServiceName != nil {
		c.SetServiceName(*req.ServiceName)
	}
	if req.Method != nil {
		c.SetMethod(req.Method.String())
	}
	if req.MethodName != nil {
		c.SetMethodName(*req.MethodName)
	}
	if req.Path != nil {
		c.SetPath(*req.Path)
	}
	if req.PathPrefix != nil {
		c.SetPathPrefix(*req.PathPrefix)
	}
	if req.Domains != nil {
		c.SetDomains(*req.Domains)
	}
	if req.Exported != nil {
		c.SetExported(*req.Exported)
	}
	if req.Depracated != nil {
		c.SetDepracated(*req.Depracated)
	}
	return c
}

func UpdateSet(u *ent.APIUpdate, req *Req) *ent.APIUpdate {
	if req.Protocol != nil {
		u.SetProtocol(req.Protocol.String())
	}
	if req.ServiceName != nil {
		u.SetServiceName(*req.ServiceName)
	}
	if req.Method != nil {
		u.SetMethod(req.Method.String())
	}
	if req.MethodName != nil {
		u.SetMethodName(*req.MethodName)
	}
	if req.Path != nil {
		u.SetPath(*req.Path)
	}
	if req.PathPrefix != nil {
		u.SetPathPrefix(*req.PathPrefix)
	}
	if req.Domains != nil {
		u.SetDomains(*req.Domains)
	}
	if req.Exported != nil {
		u.SetExported(*req.Exported)
	}
	if req.Depracated != nil {
		u.SetDepracated(*req.Depracated)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	Protocol    *cruder.Cond
	ServiceName *cruder.Cond
	Method      *cruder.Cond
	Path        *cruder.Cond
	Exported    *cruder.Cond
	Depracated  *cruder.Cond
	IDs         *cruder.Cond
}

func SetQueryConds(q *ent.APIQuery, conds *Conds) (*ent.APIQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(api.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.Protocol != nil {
		protocol, ok := conds.Protocol.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid protocol")
		}
		switch conds.Protocol.Op {
		case cruder.EQ:
			q.Where(api.Protocol(protocol))
		default:
			return nil, fmt.Errorf("invalid protocol field")
		}
	}
	if conds.ServiceName != nil {
		switch conds.ServiceName.Op {
		case cruder.EQ:
			q.Where(api.ServiceName(conds.ServiceName.Op))
		default:
			return nil, fmt.Errorf("invalid service name field")
		}
	}
	if conds.Method != nil {
		method, ok := conds.Method.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid method")
		}
		switch conds.Method.Op {
		case cruder.EQ:
			q.Where(api.Method(method))
		default:
			return nil, fmt.Errorf("invalid method field")
		}
	}
	if conds.Path != nil {
		switch conds.Path.Op {
		case cruder.EQ:
			q.Where(api.Path(conds.Path.Op))
		default:
			return nil, fmt.Errorf("invalid path field")
		}
	}
	if conds.Exported != nil {
		exported, ok := conds.Exported.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid exported")
		}
		switch conds.Exported.Op {
		case cruder.EQ:
			q.Where(api.Exported(exported))
		default:
			return nil, fmt.Errorf("invalid exported field")
		}
	}
	if conds.Depracated != nil {
		deprecated, ok := conds.Depracated.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid deprecated")
		}
		switch conds.Depracated.Op {
		case cruder.EQ:
			q.Where(api.Depracated(deprecated))
		default:
			return nil, fmt.Errorf("invalid deprecated field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(api.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid api ids filed")
		}
	}
	return q, nil
}
