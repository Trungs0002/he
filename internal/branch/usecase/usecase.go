package usecase

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/pkg/util"
)

func (uc implUsecase) Create(ctx context.Context, sc models.Scope, input branch.CreateInput) (models.Branch, error) {
	branch, err := uc.repo.Create(ctx, sc, branch.CreateOptions{
		Name:  input.Name,
		Code:  util.BuildCode(input.Name),
		Alias: util.BuildAlias(input.Name),
	})
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Create.Create: %v", err)
		return models.Branch{}, err
	}

	return branch, nil
}
