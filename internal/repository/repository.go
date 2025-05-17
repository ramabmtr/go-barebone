package repository

import (
	"github.com/ramabmtr/go-barebone/internal/config"
	"github.com/ramabmtr/go-barebone/internal/entity"
)

type Repository struct {
	Dummy entity.IDummy
}

func InitRepository() *Repository {
	db := config.GetDBConn()
	return &Repository{
		Dummy: NewDummyRepository(db),
	}
}
