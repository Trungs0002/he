package branch

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Branch, error)
	GetCurrrentRegion(ctx context.Context, sc models.Scope) (models.Region, error)
	GetAll(ctx context.Context, sc models.Scope, branch models.Branch) ([]models.Branch, error)
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Branch, error)
	Delete(ctx context.Context, sc models.Scope) error
	Get(ctx context.Context, sc models.Scope) (models.Branch, error)
}