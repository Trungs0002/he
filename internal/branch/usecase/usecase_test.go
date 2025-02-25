package usecase

import (
	"testing"

	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type mockDeps struct {
	repo *branch.MockRepository
}

func initUseCase(t *testing.T) (branch.Usecase, mockDeps) {
	repo := branch.NewMockRepository(t)
	l := log.InitializeZapLogger(log.NewTestZapConfig())

	return New(l, repo), mockDeps{
		repo: repo,
	}
}
