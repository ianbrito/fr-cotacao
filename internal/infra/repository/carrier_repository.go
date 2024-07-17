package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	"github.com/ianbrito/fr-cotacao/internal/infra/db"
	"time"
)

type CarrierRepository interface {
	Save(carrier *entity.Carrier) (*entity.Carrier, error)
	GetByReference(reference int) (*entity.Carrier, error)
	CreateOrFirst(carrier *entity.Carrier) (*entity.Carrier, error)
}

type SQLCarrierRepository struct {
	Queries *db.Queries
	Ctx     context.Context
}

func NewSQLCarrierRepository(ctx context.Context, queries *db.Queries) *SQLCarrierRepository {
	if queries == nil {
		conn := db.GetConnection()
		queries = db.New(conn)
	}
	return &SQLCarrierRepository{
		Queries: queries,
		Ctx:     ctx,
	}
}

func (r *SQLCarrierRepository) Save(carrier *entity.Carrier) (*entity.Carrier, error) {
	result, err := r.Queries.CreateCarrier(r.Ctx, db.CreateCarrierParams{
		Reference:        carrier.Reference,
		Name:             carrier.Name,
		RegisteredNumber: carrier.RegisteredNumber,
		StateInscription: carrier.StateInscription,
		LogoUrl:          carrier.LogoUrl,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	if err != nil {
		return nil, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return entity.NewCarrier(
		carrier.Reference,
		carrier.Name,
		carrier.RegisteredNumber,
		carrier.StateInscription,
		carrier.LogoUrl,
	), nil
}

func (r *SQLCarrierRepository) GetByReference(reference int64) (*entity.Carrier, error) {
	carrier, err := r.Queries.GetCarrierByID(r.Ctx, reference)
	if err != nil {
		return nil, err
	}

	return entity.NewCarrier(
		carrier.Reference,
		carrier.Name,
		carrier.RegisteredNumber,
		carrier.StateInscription,
		carrier.LogoUrl,
	), nil
}

func (r *SQLCarrierRepository) CreateOrFirst(carrier *entity.Carrier) (*entity.Carrier, error) {
	c, err := r.GetByReference(carrier.Reference)
	if errors.Is(err, sql.ErrNoRows) {
		return r.Save(carrier)
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}
