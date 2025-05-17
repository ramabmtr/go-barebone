package service

import (
	"context"

	"github.com/ramabmtr/go-barebone/internal/entity"
)

type Dummy struct {
	iDummy entity.IDummy
}

func NewDummyService(iDummy entity.IDummy) *Dummy {
	return &Dummy{
		iDummy: iDummy,
	}
}

func (c *Dummy) Dummy(ctx context.Context, caller string) error {
	return c.iDummy.Dummy(ctx, caller)
}
