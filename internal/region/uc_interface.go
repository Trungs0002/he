package region

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	// Create creates a new todo
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Region, error)
	GetAll(ctx context.Context, sc models.Scope) ([]models.Region, error)
	Update(ctx context.Context, sc models.Scope, input UpdateInput) (models.Region, error)
	Delete(ctx context.Context, sc models.Scope) error
	Get(ctx context.Context, sc models.Scope) (models.Region, error)
}