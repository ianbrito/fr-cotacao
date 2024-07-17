package models

import (
	"encoding/json"
	"log"
)

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type             int    `json:"type"`
	RegisteredNumber string `json:"registered_number"`
	Country          string `json:"country"`
	ZipCode          uint32 `json:"zipcode"`
}

type Volume struct {
	Category      string  `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	UnitaryPrice  float64 `json:"unitary_price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

type Dispatcher struct {
	RegisteredNumber string    `json:"registered_number"`
	ZipCode          uint32    `json:"zipcode"`
	Volumes          []*Volume `json:"volumes"`
}

type QuoteRequest struct {
	Shipper        Shipper       `json:"shipper"`
	Recipient      Recipient     `json:"recipient"`
	Dispatchers    []*Dispatcher `json:"dispatchers"`
	SimulationType []int         `json:"simulation_type"`
}

func NewQuoteRequest(shipper Shipper, recipient Recipient, dispatchers []*Dispatcher) *QuoteRequest {
	return &QuoteRequest{
		Shipper:     shipper,
		Recipient:   recipient,
		Dispatchers: dispatchers,
	}
}

func (qr *QuoteRequest) AddDispatcher(dispatcher Dispatcher) {
	qr.Dispatchers = append(qr.Dispatchers, &dispatcher)
}

func (qr *QuoteRequest) ParseToJSON() []byte {
	data, err := json.Marshal(qr)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
