package usecase

import (
	"gitlab.com/gma-vietnam/tanca-event/internal/region"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type implUsecase struct {
	l    log.Logger
	repo region.Repository
}

func New(l log.Logger, repo region.Repository) region.Usecase {
	return &implUsecase{
		l:    l,
		repo: repo,
	}
}