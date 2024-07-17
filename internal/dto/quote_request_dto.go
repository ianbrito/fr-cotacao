package dto

import (
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	"strconv"
)

type AddressRequest struct {
	ZipCode string `json:"zipcode" validate:"required"`
}

type RecipientRequest struct {
	Address *AddressRequest `json:"address" validate:"required"`
}

type VolumeRequest struct {
	Category      int     `json:"category" validate:"required"`
	Amount        int     `json:"amount" validate:"required"`
	UnitaryWeight float64 `json:"unitary_weight" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	Sku           string  `json:"sku" validate:"required"`
	Height        float64 `json:"height" validate:"required"`
	Width         float64 `json:"width" validate:"required"`
	Length        float64 `json:"length" validate:"required"`
}

type QuoteRequest struct {
	Recipient *RecipientRequest `json:"recipient" validate:"required"`
	Volumes   []*VolumeRequest  `json:"volumes" validate:"required,dive"`
}

func (qr *QuoteRequest) ToEntity() (*entity.Shipper, *entity.Recipient, *entity.Dispatcher) {
	shipper := entity.NewShipper()
	recipient := entity.NewRecipient(qr.Recipient.Address.ZipCode)
	dispatcher := entity.NewDispatcher()
	for _, v := range qr.Volumes {
		dispatcher.AddVolume(entity.NewVolume(
			strconv.Itoa(v.Category),
			v.Amount,
			v.UnitaryWeight,
			v.Price,
			v.Sku,
			v.Height,
			v.Width,
			v.Length,
		))
	}
	return shipper, recipient, dispatcher
}
