package repository

import (
	"context"

	"github.com/ramabmtr/go-barebone/internal/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type dummy struct {
	db *gorm.DB
}

func NewDummyRepository(db *gorm.DB) entity.IDummy {
	return &dummy{
		db: db,
	}
}

func (p *dummy) Dummy(ctx context.Context, caller string) error {
	log.Info().Msgf("caller: %s", caller)
	return nil
}
