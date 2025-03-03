package department

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Department, error)
	GetCurrrentBranch(ctx context.Context, sc models.Scope) (models.Branch, error)
	GetAll(ctx context.Context, sc models.Scope, branch models.Branch) ([]models.Department, error)
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Department, error)
	Delete(ctx context.Context, sc models.Scope) error
	Get(ctx context.Context, sc models.Scope) (models.Department, error)
}