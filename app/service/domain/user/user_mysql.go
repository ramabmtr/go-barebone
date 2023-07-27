package user

import (
	"context"

	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"gorm.io/gorm"
)

type userMySQL struct {
	db *gorm.DB
}

func NewUserDomainMySQL(db *gorm.DB) User {
	return &userMySQL{
		db: db,
	}
}

func (ud *userMySQL) Create(ctx context.Context, u *entity.User) error {
	return ud.db.
		WithContext(ctx).
		Model(u).
		Create(u).Error
}

func (ud *userMySQL) Get(ctx context.Context, u *entity.User) error {
	err := ud.db.
		WithContext(ctx).
		Model(u).
		Take(u).Error
	if err == gorm.ErrRecordNotFound {
		return errors.ErrDataNotFound
	}
	if err != nil {
		return err
	}

	return nil
}
