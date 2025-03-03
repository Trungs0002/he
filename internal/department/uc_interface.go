package department

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

//go:generate mockery --name=Usecase
type Usecase interface {
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Department, error)
	GetAll(ctx context.Context, sc models.Scope) ([]models.Department, error)
	Update(ctx context.Context, sc models.Scope, input UpdateInput) (models.Department, error)
	Delete(ctx context.Context, sc models.Scope) error
	Get(ctx context.Context, sc models.Scope) (models.Department, error)
}