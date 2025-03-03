package mongo

import (
	"time"

	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/internal/region"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"
)

type implRepository struct {
	l        log.Logger
	db       mongo.Database
	regionUC region.Usecase
	clock    func() time.Time
}

func NewRepository(l log.Logger, db mongo.Database, regionUC region.Usecase) branch.Repository {
	return &implRepository{
		l:        l,
		db:       db,
		regionUC: regionUC,
		clock:    time.Now,
	}
}