package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	// ID          primitive.ObjectID `bson:"_id"`
	// Name        string             `bson:"name"`
	// ShopID      primitive.ObjectID `bson:"shop_id"`
	// RegionID    primitive.ObjectID `bson:"region_id"`
	// BranchID    primitive.ObjectID `bson:"branch_id"`
	// DepartmentID *primitive.ObjectID `bson:"department_id,omitempty"` 
	// CreatedAt   time.Time          `bson:"created_at"`
	// UpdatedAt   time.Time          `bson:"updated_at"`
	// DeletedAt   *time.Time         `bson:"deleted_at,omitempty"`

	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`

	ShopID       primitive.ObjectID  `bson:"shop_id"`
	RegionID     *primitive.ObjectID `bson:"region_id,omitempty"`
	BranchID     *primitive.ObjectID `bson:"branch_id,omitempty"`
	DepartmentID *primitive.ObjectID `bson:"department_id,omitempty"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}