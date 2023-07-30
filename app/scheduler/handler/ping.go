package handler

import (
	"context"

	"github.com/rs/zerolog/log"
)

func Ping(ctx context.Context) error {
	log.Info().Ctx(ctx).Msg("schedule pong")
	return nil
}
