package repository

import (
	"context"
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	"github.com/ianbrito/fr-cotacao/internal/infra/db"
	"time"
)

type DispatcherRepository interface {
	Save(dispatcher *entity.Dispatcher)
}

type SQLDispatcherRepository struct {
	Queries         *db.Queries
	Ctx             context.Context
	OfferRepository *SQLOfferRepository
}

func NewSQLDispatcherRepository(ctx context.Context) *SQLDispatcherRepository {
	return &SQLDispatcherRepository{
		Queries:         db.New(DB),
		Ctx:             ctx,
		OfferRepository: NewSQLOfferRepository(ctx),
	}
}

func (r *SQLDispatcherRepository) Save(dispatcher *entity.Dispatcher) (*entity.Dispatcher, error) {
	_, err := r.Queries.CreateDispatcher(r.Ctx, db.CreateDispatcherParams{
		ID:                         dispatcher.ID,
		RequestID:                  dispatcher.RequestID,
		RegisteredNumberDispatcher: dispatcher.RegisteredNumberDispatcher,
		RegisteredNumberShipper:    dispatcher.RegisteredNumberShipper,
		ZipcodeOrigin:              int32(dispatcher.ZipCode),
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	})

	if err != nil {
		return nil, err
	}

	for _, offer := range dispatcher.Offers {
		_, err := r.OfferRepository.Save(offer, dispatcher.ID)
		if err != nil {
			return nil, err
		}
	}

	return dispatcher, nil
}
