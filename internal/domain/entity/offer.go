package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Offer struct {
	ID                           string
	Offer                        int32
	SimulationType               int32
	Carrier                      *Carrier
	Service                      string
	ServiceCode                  string
	ServiceDescription           string
	Identifier                   string
	DeliveryNote                 string
	HomeDelivery                 bool
	CarrierNeedsToReturnToSender bool
	Expiration                   time.Time
	CostPrice                    float64
	FinalPrice                   float64
	Modal                        string
	Weights                      *Weights
	DeliveryTime                 *DeliveryTime
	OriginalDeliveryTime         *DeliveryTime
}

type Weights struct {
	Real  float64  `json:"real"`
	Cubed *float64 `json:"cubed"`
	Used  float64  `json:"used"`
}

type DeliveryTime struct {
	Days          int    `json:"days"`
	Hours         int    `json:"hours"`
	Minutes       int    `json:"minutes"`
	EstimatedDate string `json:"estimated_date"`
}

func NewOffer(
	offer int32,
	simulationType int32,
	carrier *Carrier,
	service string,
	serviceCode string,
	serviceDescription string,
	identifier string,
	deliveryNote string,
	homeDelivery bool,
	carrierNeedsToReturnToSender bool,
	expiration string,
	costPrice float64,
	finalPrice float64,
	modal string,
	weights *Weights,
	deliveryTime *DeliveryTime,
	originalDeliveryTime *DeliveryTime,
) *Offer {
	parsedExpiration, err := time.Parse(time.RFC3339Nano, expiration)
	if err != nil {
		fmt.Println("Erro ao converter a string para time.Time:", err)
	}

	return &Offer{
		ID:                           uuid.NewString(),
		Offer:                        offer,
		SimulationType:               simulationType,
		Carrier:                      carrier,
		Service:                      service,
		ServiceCode:                  serviceCode,
		ServiceDescription:           serviceDescription,
		Identifier:                   identifier,
		DeliveryNote:                 deliveryNote,
		HomeDelivery:                 homeDelivery,
		CarrierNeedsToReturnToSender: carrierNeedsToReturnToSender,
		Expiration:                   parsedExpiration,
		CostPrice:                    costPrice,
		FinalPrice:                   finalPrice,
		Modal:                        modal,
		Weights:                      weights,
		DeliveryTime:                 deliveryTime,
		OriginalDeliveryTime:         originalDeliveryTime,
	}
}
