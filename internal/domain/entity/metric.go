package entity

type Metric struct {
	CheaperShipping       float64
	MostExpensiveShipping float64
	CarrierMetrics        []*CarrierMetric
}

type CarrierMetric struct {
	Name           string
	Total          int64
	FinalPriceSum  float64
	FinalPriceMean float64
}
