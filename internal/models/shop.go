package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	// ID        primitive.ObjectID `bson:"_id"`
	// Name      string             `bson:"name"`
	// Regions   []Region           `bson:"regions"`
	// CreatedAt time.Time          `bson:"created_at"`
	// UpdatedAt time.Time          `bson:"updated_at"`
	// DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	Region      []primitive.ObjectID `bson:"regions"`
	CreatedAt   time.Time            `bson:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at"`
	DeletedAt   *time.Time           `bson:"deleted_at,omitempty"`
}