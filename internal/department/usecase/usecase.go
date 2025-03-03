package usecase

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/department"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
)

func (uc implUsecase) Create(ctx context.Context, sc models.Scope, input department.CreateInput) (models.Department, error) {
	department, err := uc.repo.Create(ctx, sc, department.CreateOptions{
		Name: input.Name,
	})
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Create.Create: %v", err)
		return models.Department{}, err
	}

	return department, nil
}
func (uc implUsecase) GetAll(ctx context.Context, sc models.Scope) ([]models.Department, error) {
	branch, err := uc.repo.GetCurrrentBranch(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.GetListDepart.GetCurrrentDepartment: %v", err)
		return nil, ErrBranchNotFound
	}
	departmentes, err := uc.repo.GetAll(ctx, sc, branch)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.GetListDepart.GetListDepart: %v", err)
		return nil, err
	}

	return departmentes, nil
}

func (uc implUsecase) Update(ctx context.Context, sc models.Scope, input department.UpdateInput) (models.Department, error) {
	department, err := uc.repo.Update(ctx, sc, department.UpdateOptions{
		ID:   sc.DepartmentID,
		Name: input.Name,
	})
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Update.Update: %v", err)
		return models.Department{}, ErrDepartmentNotFound
	}

	return department, nil
}

func (uc implUsecase) Delete(ctx context.Context, sc models.Scope) error {
	department, err := uc.repo.Get(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Delete.Get: %v", err)
		return ErrDepartmentNotFound
	}
	if department.Branches != nil {
		return ErrDepartmentHasBranch
	}
	err = uc.repo.Delete(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Delete.Delete: %v", err)
		return ErrDepartmentNotFound
	}

	return nil
}

func (uc implUsecase) Get(ctx context.Context, sc models.Scope) (models.Department, error) {
	department, err := uc.repo.Get(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "department.usecase.Get.Get: %v", err)
		return models.Department{}, err
	}

	return department, nil
}