package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Region struct {
	// ID        primitive.ObjectID `bson:"_id"`
	// Name      string             `bson:"name"`
	// Branches  []Branch           `bson:"branches"`
	// ShopID    primitive.ObjectID `bson:"shop_id"`
	// CreatedAt time.Time          `bson:"created_at"`
	// UpdatedAt time.Time          `bson:"updated_at"`
	// DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
	ID        primitive.ObjectID   `bson:"_id"`
	Name      string               `bson:"name"`
	Branches  []primitive.ObjectID `bson:"branches"`
	CreatedAt time.Time            `bson:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at"`
	DeletedAt *time.Time           `bson:"deleted_at,omitempty"`
}