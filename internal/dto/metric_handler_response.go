package dto

import "net/http"

type MetricResponse struct {
	CheaperShipping       float64                  `json:"frete_mais_barato"`
	MostExpensiveShipping float64                  `json:"frete_mais_caro"`
	CarrierMetrics        []*CarrierMetricResponse `json:"carrier"`
}

type CarrierMetricResponse struct {
	Name           string  `json:"transportadora"`
	Total          int64   `json:"resultados"`
	FinalPriceSum  float64 `json:"total_preco_frete"`
	FinalPriceMean float64 `json:"media_preco_frete"`
}

func (qr *MetricResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
