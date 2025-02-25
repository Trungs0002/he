package mongo

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
)

const (
	branchCollection = "branches"
)

func (repo implRepository) getBranchCollection() mongo.Collection {
	return repo.db.Collection(branchCollection)
}

func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts branch.CreateOptions) (models.Branch, error) {
	col := repo.getBranchCollection()
	now := repo.clock()

	branch := models.Branch{
		ID:        repo.db.NewObjectID(),
		Name:      opts.Name,
		Alias:     opts.Alias,
		Code:      opts.Code,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := col.InsertOne(ctx, branch)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Create.InsertOne: %v", err)
		return models.Branch{}, err
	}

	return branch, err
}
