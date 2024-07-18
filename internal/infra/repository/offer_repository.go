package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	"github.com/ianbrito/fr-cotacao/internal/infra/db"
	"time"
)

type OfferRepository interface {
	Save(offer *entity.Offer, dispatcherID string)
}

type SQLOfferRepository struct {
	Queries           *db.Queries
	Ctx               context.Context
	CarrierRepository *SQLCarrierRepository
}

func NewSQLOfferRepository(ctx context.Context) *SQLOfferRepository {
	return &SQLOfferRepository{
		Queries:           db.New(DB),
		Ctx:               ctx,
		CarrierRepository: NewSQLCarrierRepository(ctx),
	}
}

func (r *SQLOfferRepository) Save(offer *entity.Offer, dispatcherID string) (string, error) {
	carrier, err := r.CarrierRepository.CreateOrFirst(offer.Carrier)
	if err != nil {
		panic(err)
	}

	weights, err := json.Marshal(offer.Weights)
	if err != nil {
		panic(err)
	}

	deliveryTime, err := json.Marshal(offer.DeliveryTime)
	if err != nil {
		panic(err)
	}

	originalDeliveryTime, err := json.Marshal(offer.OriginalDeliveryTime)
	if err != nil {
		panic(err)
	}

	result, err := r.Queries.CreateOffer(r.Ctx, db.CreateOfferParams{
		ID:                           offer.ID,
		DispatcherID:                 dispatcherID,
		Offer:                        offer.Offer,
		SimulationType:               offer.SimulationType,
		CarrierID:                    carrier.Reference,
		Service:                      offer.Service,
		ServiceCode:                  toNullString(offer.ServiceCode),
		ServiceDescription:           toNullString(offer.ServiceDescription),
		Identifier:                   toNullString(offer.Identifier),
		DeliveryNote:                 toNullString(offer.DeliveryNote),
		HomeDelivery:                 sql.NullBool{Bool: offer.HomeDelivery, Valid: true},
		CarrierNeedsToReturnToSender: sql.NullBool{Bool: offer.CarrierNeedsToReturnToSender, Valid: true},
		Expiration:                   offer.Expiration,
		CostPrice:                    offer.CostPrice,
		FinalPrice:                   offer.FinalPrice,
		Modal:                        toNullString(offer.Modal),
		Weights:                      weights,
		DeliveryTime:                 deliveryTime,
		OriginalDeliveryTime:         originalDeliveryTime,
		CreatedAt:                    time.Now(),
		UpdatedAt:                    time.Now(),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return "", err
	}

	return offer.ID, nil
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}
