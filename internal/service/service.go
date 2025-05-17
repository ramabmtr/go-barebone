package service

import "github.com/ramabmtr/go-barebone/internal/repository"

type Service struct {
	Dummy *Dummy
}

func InitService(repo *repository.Repository) *Service {
	return &Service{
		Dummy: NewDummyService(repo.Dummy),
	}
}
