package mongo

import (
	"context"

	"gitlab.com/gma-vietnam/tanca-event/internal/department"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	departmentCollection = "departments"
	branchCollection     = "branchs"
)

func (repo implRepository) getDepartmentCollection() mongo.Collection {
	return repo.db.Collection(departmentCollection)
}
func (repo implRepository) getBranchCollection() mongo.Collection {
	return repo.db.Collection(branchCollection)
}

func (repo implRepository) Create(ctx context.Context, sc models.Scope, opts department.CreateOptions) (models.Department, error) {
	col := repo.getDepartmentCollection()
	now := repo.clock()

	department := models.Department{
		ID:        repo.db.NewObjectID(),
		Name:      opts.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := col.InsertOne(ctx, department)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Create.InsertOne: %v", err)
		return models.Department{}, err
	}

	return department, err
}
func (repo implRepository) GetCurrrentBranch(ctx context.Context, sc models.Scope) (models.Branch, error) {
	col := repo.getBranchCollection()
	var branch models.Branch
	branchID, err := primitive.ObjectIDFromHex(sc.BranchID)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.GetCurrrentDepartment.ObjectIDFromHex: %v", err)
		return models.Branch{}, err
	}
	err = col.FindOne(ctx, bson.M{"_id": branchID}).Decode(&branch)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.GetCurrrentDepartment.FindOne: %v", err)
		return models.Branch{}, err
	}
	return branch, nil
}
func (repo implRepository) GetAll(ctx context.Context, sc models.Scope, branch models.Branch) ([]models.Department, error) {
	col := repo.getDepartmentCollection()
	var departmentes []models.Department
	cur, err := col.Find(ctx, bson.M{"_id": bson.M{"$in": branch.Department}})
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.GetListDepart.Find: %v", err)
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var department models.Department
		err := cur.Decode(&department)
		if err != nil {
			repo.l.Errorf(ctx, "event.department.mongo.GetListDepart.Decode: %v", err)
			return nil, err
		}
		departmentes = append(departmentes, department)
	}
	return departmentes, nil
}
func (repo implRepository) Update(ctx context.Context, sc models.Scope, opts department.UpdateOptions) (models.Department, error) {
	col := repo.getDepartmentCollection()
	now := repo.clock()
	departmentID, err := primitive.ObjectIDFromHex(sc.DepartmentID)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Update.ObjectIDFromHex: %v", sc.DepartmentID)
		return models.Department{}, err
	}
	update := bson.M{
		"$set": bson.M{
			"name":       opts.Name,
			"updated_at": now,
		},
	}
	var department models.Department
	err = col.FindOne(ctx, bson.M{"_id": departmentID}).Decode(&department)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Update.FindOne: %v", err)
		return models.Department{}, err
	}
	_, err = col.UpdateOne(ctx, bson.M{"_id": departmentID}, update)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Update.FindOneAndUpdate: %v", err)
		return models.Department{}, err
	}
	err = col.FindOne(ctx, bson.M{"_id": departmentID}).Decode(&department)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Update.Result: %v", err)
		return models.Department{}, err
	}
	return department, nil
}

func (repo implRepository) Delete(ctx context.Context, sc models.Scope) error {
	col := repo.getDepartmentCollection()
	departmentIDObj, err := primitive.ObjectIDFromHex(sc.DepartmentID)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Delete.ObjectIDFromHex: %v", err)
		return err
	}
	_, err = col.DeleteOne(ctx, bson.M{"_id": departmentIDObj})
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Delete.DeleteOne: %v", err)
		return err
	}
	return nil
}

func (repo implRepository) Get(ctx context.Context, sc models.Scope) (models.Department, error) {
	col := repo.getDepartmentCollection()
	departmentID, err := primitive.ObjectIDFromHex(sc.DepartmentID)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Get.ObjectIDFromHex: %v", err)
		return models.Department{}, err
	}
	var department models.Department
	repo.l.Infof(ctx, "departmentID: %v", departmentID)
	err = col.FindOne(ctx, bson.M{"_id": departmentID}).Decode(&department)
	if err != nil {
		repo.l.Errorf(ctx, "event.department.mongo.Get.FindOne: %v", err)
		return models.Department{}, err
	}
	return department, nil
}