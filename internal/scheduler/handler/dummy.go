package handler

import (
	"context"

	"github.com/ramabmtr/go-barebone/internal/service"
)

type Dummy struct {
	dummySvc *service.Dummy
}

func NewDummyHandler(dummySvc *service.Dummy) *Dummy {
	return &Dummy{
		dummySvc: dummySvc,
	}
}

func (h *Dummy) Dummy(ctx context.Context) error {
	err := h.dummySvc.Dummy(ctx, "scheduler")
	if err != nil {
		return err
	}
	return nil
}
