package usercode

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Handler struct {
	Prefix      *string
	AppID       *string
	Account     *string
	AccountType *string
	UsedFor     *string
	VCode       *string
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

func WithPrefix(prefix *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if prefix == nil {
			return fmt.Errorf("prefix is empty")
		}
		h.Prefix = prefix
		return nil
	}
}

func WithAppID(appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_, err := uuid.Parse(*appID)
		if err != nil {
			return err
		}
		h.AppID = appID
		return nil
	}
}

func WithAccount(account *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if account == nil {
			return fmt.Errorf("account is empty")
		}
		h.Account = account
		return nil
	}
}

func WithAccountType(_type *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return fmt.Errorf("account type is empty")
		}
		h.AccountType = _type
		return nil
	}
}

func WithUsedFor(usedFor *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if usedFor == nil {
			return fmt.Errorf("UsedFor %v is invalid", *usedFor)
		}
		h.UsedFor = usedFor
		return nil
	}
}

func WithCode(code *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if code == nil {
			return fmt.Errorf("code is empty")
		}
		h.VCode = code
		return nil
	}
}
