package dto

import (
	"encoding/json"
	"net/http"
)

type QuoteResponse struct {
	Carrier []*CarrierResponse `json:"carrier"`
}

type CarrierResponse struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}

func NewQuoteResponse() *QuoteResponse {
	return &QuoteResponse{}
}

func (qr *QuoteResponse) AddCarrier(carrier *CarrierResponse) {
	qr.Carrier = append(qr.Carrier, carrier)
}

func (qr *QuoteResponse) ParseToJson() ([]byte, error) {
	data, err := json.Marshal(qr)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (qr *QuoteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
