package branch

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	// Create creates a new todo
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Branch, error)
}
