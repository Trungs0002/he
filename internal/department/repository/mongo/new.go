package mongo

import (
	"time"

	"gitlab.com/gma-vietnam/tanca-event/internal/department"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
)

type implRepository struct {
	l     log.Logger
	db    mongo.Database
	clock func() time.Time
}

func NewRepository(l log.Logger, db mongo.Database) department.Repository {
	return &implRepository{
		l:     l,
		db:    db,
		clock: time.Now,
	}
}