package usecase

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/internal/region"
)

func (uc implUsecase) Create(ctx context.Context, sc models.Scope, input region.CreateInput) (models.Region, error) {
	region, err := uc.repo.Create(ctx, sc, region.CreateOptions{
		Name: input.Name,
	})
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Create.Create: %v", err)
		return models.Region{}, err
	}

	return region, nil
}
func (uc implUsecase) GetAll(ctx context.Context, sc models.Scope) ([]models.Region, error) {
	shop, err := uc.repo.GetCurrrentShop(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.GetListDepart.GetCurrrentRegion: %v", err)
		return nil, ErrShopNotFound
	}
	regiones, err := uc.repo.GetAll(ctx, sc, shop)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.GetListDepart.GetListDepart: %v", err)
		return nil, err
	}

	return regiones, nil
}

func (uc implUsecase) Update(ctx context.Context, sc models.Scope, input region.UpdateInput) (models.Region, error) {
	region, err := uc.repo.Update(ctx, sc, region.UpdateOptions{
		ID:   sc.RegionID,
		Name: input.Name,
	})
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Update.Update: %v", err)
		return models.Region{}, ErrRegionNotFound
	}

	return region, nil
}

func (uc implUsecase) Delete(ctx context.Context, sc models.Scope) error {
	region, err := uc.repo.Get(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Delete.Get: %v", err)
		return ErrRegionNotFound
	}
	if region.Branches != nil {
		return ErrRegionHasBranch
	}
	err = uc.repo.Delete(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Delete.Delete: %v", err)
		return ErrRegionNotFound
	}

	return nil
}

func (uc implUsecase) Get(ctx context.Context, sc models.Scope) (models.Region, error) {
	region, err := uc.repo.Get(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "region.usecase.Get.Get: %v", err)
		return models.Region{}, err
	}

	return region, nil
}