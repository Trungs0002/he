package usecase

import (
	"gitlab.com/gma-vietnam/tanca-event/internal/department"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type implUsecase struct {
	l    log.Logger
	repo department.Repository
}

func New(l log.Logger, repo department.Repository) department.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}