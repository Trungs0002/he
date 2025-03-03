package region

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

// Repository is the interface for todo repository
//
//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Region, error)
	GetCurrrentShop(ctx context.Context, sc models.Scope) (models.Shop, error)
	GetAll(ctx context.Context, sc models.Scope, shop models.Shop) ([]models.Region, error)
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Region, error)
	Delete(ctx context.Context, sc models.Scope) error
	Get(ctx context.Context, sc models.Scope) (models.Region, error)
}