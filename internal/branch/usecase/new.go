package usecase

import (
	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type implUsecase struct {
	l    log.Logger
	repo branch.Repository
}

func New(l log.Logger, repo branch.Repository) branch.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}
