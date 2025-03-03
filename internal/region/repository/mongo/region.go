package mongo

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/internal/region"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	regionCollection = "regions"
	shopCollection   = "shops"
)

func (repo implRepository) getRegionCollection() mongo.Collection {
	return repo.db.Collection(regionCollection)
}
func (repo implRepository) getShopCollection() mongo.Collection {
	return repo.db.Collection(shopCollection)
}

func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts region.CreateOptions) (models.Region, error) {
	col := repo.getRegionCollection()
	now := repo.clock()

	region := models.Region{
		ID:        repo.db.NewObjectID(),
		Name:      opts.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := col.InsertOne(ctx, region)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Create.InsertOne: %v", err)
		return models.Region{}, err
	}

	return region, err
}
func (repo implRepository) GetCurrrentShop(ctx context.Context, sc models.Scope) (models.Shop, error) {
	col := repo.getShopCollection()
	var shop models.Shop
	shopID, err := primitive.ObjectIDFromHex(sc.ShopID)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.GetCurrrentRegion.ObjectIDFromHex: %v", err)
		return models.Shop{}, err
	}
	err = col.FindOne(ctx, bson.M{"_id": shopID}).Decode(&shop)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.GetCurrrentRegion.FindOne: %v", err)
		return models.Shop{}, err
	}
	return shop, nil
}
func (repo implRepository) GetAll(ctx context.Context, sc models.Scope, shop models.Shop) ([]models.Region, error) {
	col := repo.getRegionCollection()
	var regiones []models.Region
	cur, err := col.Find(ctx, bson.M{"_id": bson.M{"$in": shop.Region}})
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.GetListDepart.Find: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var region models.Region
		err := cur.Decode(&region)
		if err != nil {
			repo.l.Errorf(ctx, "event.region.mongo.GetListDepart.Decode: %v", err)
			return nil, err
		}
		regiones = append(regiones, region)
	}
	return regiones, nil
}
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts region.UpdateOptions) (models.Region, error) {
	col := repo.getRegionCollection()
	now := repo.clock()
	regionID, err := primitive.ObjectIDFromHex(sc.RegionID)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Update.ObjectIDFromHex: %v", sc.RegionID)
		return models.Region{}, err
	}
	update := bson.M{
		"$set": bson.M{
			"name":       opts.Name,
			"updated_at": now,
		},
	}
	var region models.Region
	err = col.FindOne(ctx, bson.M{"_id": regionID}).Decode(&region)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Update.FindOne: %v", err)
		return models.Region{}, err
	}
	_, err = col.UpdateOne(ctx, bson.M{"_id": regionID}, update)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Update.FindOneAndUpdate: %v", err)
		return models.Region{}, err
	}
	err = col.FindOne(ctx, bson.M{"_id": regionID}).Decode(&region)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Update.Result: %v", err)
		return models.Region{}, err
	}
	return region, nil
}

func (repo implRepository) Delete(ctx context.Context, sc models.Scope) error {
	col := repo.getRegionCollection()
	regionIDObj, err := primitive.ObjectIDFromHex(sc.RegionID)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Delete.ObjectIDFromHex: %v", err)
		return err
	}
	_, err = col.DeleteOne(ctx, bson.M{"_id": regionIDObj})
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Delete.DeleteOne: %v", err)
		return err
	}
	return nil
}

func (repo implRepository) Get(ctx context.Context, sc models.Scope) (models.Region, error) {
	col := repo.getRegionCollection()
	regionID, err := primitive.ObjectIDFromHex(sc.RegionID)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Get.ObjectIDFromHex: %v", err)
		return models.Region{}, err
	}
	var region models.Region
	repo.l.Infof(ctx, "regionID: %v", regionID)
	err = col.FindOne(ctx, bson.M{"_id": regionID}).Decode(&region)
	if err != nil {
		repo.l.Errorf(ctx, "event.region.mongo.Get.FindOne: %v", err)
		return models.Region{}, err
	}
	return region, nil
}