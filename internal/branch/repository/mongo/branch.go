package mongo

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	branchCollection = "branches"
	regionCollection = "regions"
)

func (repo implRepository) getBranchCollection() mongo.Collection {
	return repo.db.Collection(branchCollection)
}
func (repo implRepository) getRegionCollection() mongo.Collection {
	return repo.db.Collection(regionCollection)
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
func (repo implRepository) GetCurrrentRegion(ctx context.Context, sc models.Scope) (models.Region, error) {
	col := repo.getRegionCollection()
	var region models.Region
	regionID, err := primitive.ObjectIDFromHex(sc.RegionID)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.GetCurrrentRegion.ObjectIDFromHex: %v", err)
		return models.Region{}, err
	}
	err = col.FindOne(ctx, bson.M{"_id": regionID}).Decode(&region)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.GetCurrrentRegion.FindOne: %v", err)
		return models.Region{}, err
	}
	return region, nil
}
func (repo implRepository) GetAll(ctx context.Context, sc models.Scope, region models.Region) ([]models.Branch, error) {
	col := repo.getBranchCollection()
	var branches []models.Branch
	cur, err := col.Find(ctx, bson.M{"_id": bson.M{"$in": region.Branches}})
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.GetListDepart.Find: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var branch models.Branch
		err := cur.Decode(&branch)
		if err != nil {
			repo.l.Errorf(ctx, "event.branch.mongo.GetListDepart.Decode: %v", err)
			return nil, err
		}
		branches = append(branches, branch)
	}
	return branches, nil
}
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts branch.UpdateOptions) (models.Branch, error) {
	col := repo.getBranchCollection()
	now := repo.clock()
	branchID, err := primitive.ObjectIDFromHex(sc.BranchID)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Update.ObjectIDFromHex: %v", sc.BranchID)
		return models.Branch{}, err
	}
	update := bson.M{
		"$set": bson.M{
			"name":       opts.Name,
			"alias":      opts.Alias,
			"code":       opts.Code,
			"updated_at": now,
		},
	}
	var branch models.Branch
	err = col.FindOne(ctx, bson.M{"_id": branchID}).Decode(&branch)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Update.FindOne: %v", err)
		return models.Branch{}, err
	}
	_, err = col.UpdateOne(ctx, bson.M{"_id": branchID}, update)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Update.FindOneAndUpdate: %v", err)
		return models.Branch{}, err
	}
	err = col.FindOne(ctx, bson.M{"_id": branchID}).Decode(&branch)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Update.Result: %v", err)
		return models.Branch{}, err
	}
	return branch, nil
}

func (repo implRepository) Delete(ctx context.Context, sc models.Scope) error {
	col := repo.getBranchCollection()
	branchIDObj, err := primitive.ObjectIDFromHex(sc.BranchID)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Delete.ObjectIDFromHex: %v", err)
		return err
	}
	_, err = col.DeleteOne(ctx, bson.M{"_id": branchIDObj})
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Delete.DeleteOne: %v", err)
		return err
	}
	return nil
}

func (repo implRepository) Get(ctx context.Context, sc models.Scope) (models.Branch, error) {
	col := repo.getBranchCollection()
	branchID, err := primitive.ObjectIDFromHex(sc.BranchID)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Get.ObjectIDFromHex: %v", err)
		return models.Branch{}, err
	}
	var branch models.Branch
	err = col.FindOne(ctx, bson.M{"_id": branchID}).Decode(&branch)
	if err != nil {
		repo.l.Errorf(ctx, "event.branch.mongo.Get.FindOne: %v", err)
		return models.Branch{}, err
	}
	return branch, nil
}