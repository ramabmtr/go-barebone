package handler

import (
	"context"
	"log"
)

func Ping(_ context.Context) error {
	log.Println("schedule pong")
	return nil
}
