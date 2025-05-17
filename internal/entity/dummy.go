package entity

import "context"

type IDummy interface {
	Dummy(ctx context.Context, caller string) error
}
