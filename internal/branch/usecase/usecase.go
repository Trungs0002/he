package usecase

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

func (uc implUsecase) Create(ctx context.Context, sc models.Scope, input branch.CreateInput) (models.Branch, error) {
	branch, err := uc.repo.Create(ctx, sc, branch.CreateOptions{
		Name:  input.Name,
		Code:  input.Code,
		Alias: input.Alias,
	})
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Create.Create: %v", err)
		return models.Branch{}, err
	}

	return branch, nil
}

func (uc implUsecase) GetAll(ctx context.Context, sc models.Scope) ([]models.Branch, error) {
	region, err := uc.repo.GetCurrrentRegion(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.GetListDepart.GetCurrrentRegion: %v", err)
		return nil, ErrRegionNotFound
	}
	branches, err := uc.repo.GetAll(ctx, sc, region)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.GetListDepart.GetListDepart: %v", err)
		return nil, err
	}

	return branches, nil
}

func (uc implUsecase) Update(ctx context.Context, sc models.Scope, input branch.UpdateInput) (models.Branch, error) {
	branch, err := uc.repo.Update(ctx, sc, branch.UpdateOptions{
		ID:    sc.BranchID,
		Name:  input.Name,
		Code:  input.Code,
		Alias: input.Alias,
	})
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Update.Update: %v", err)
		return models.Branch{}, ErrBranchNotFound
	}

	return branch, nil
}

func (uc implUsecase) Delete(ctx context.Context, sc models.Scope) error {
	branch, err := uc.repo.Get(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Delete.Get: %v", err)
		return ErrBranchNotFound
	}
	if branch.Department != nil {
		return ErrBranchHasDepartment
	}
	err = uc.repo.Delete(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Delete.Delete: %v", err)
		return ErrBranchNotFound
	}

	return nil
}

func (uc implUsecase) Get(ctx context.Context, sc models.Scope) (models.Branch, error) {
	branch, err := uc.repo.Get(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "branch.usecase.Get.Get: %v", err)
		return models.Branch{}, err
	}

	return branch, nil
}