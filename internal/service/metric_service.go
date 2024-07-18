package service

import (
	"context"
	"github.com/ianbrito/fr-cotacao/internal/dto"
	"github.com/ianbrito/fr-cotacao/internal/infra/repository"
)

type MetricService struct {
	Ctx        context.Context
	Repository *repository.SQLMetricRepository
}

func NewMetricService(ctx context.Context) *MetricService {
	return &MetricService{
		Ctx:        ctx,
		Repository: repository.NewSQLMetricRepository(ctx),
	}
}

func (ms *MetricService) GetMetrics() (*dto.MetricResponse, error) {
	metric, err := ms.Repository.GetMetric()
	if err != nil {
		return nil, err
	}

	var carrierMetrics []*dto.CarrierMetricResponse
	for _, c := range metric.CarrierMetrics {
		carrierMetrics = append(carrierMetrics, &dto.CarrierMetricResponse{
			Name:           c.Name,
			Total:          c.Total,
			FinalPriceMean: c.FinalPriceMean,
			FinalPriceSum:  c.FinalPriceSum,
		})
	}

	return &dto.MetricResponse{
		CheaperShipping:       metric.CheaperShipping,
		MostExpensiveShipping: metric.MostExpensiveShipping,
		CarrierMetrics:        carrierMetrics,
	}, nil
}

func (ms *MetricService) GetMetricsWithLimit(limit int) (*dto.MetricResponse, error) {
	metric, err := ms.Repository.GetMetricWithLimit(limit)
	if err != nil {
		return nil, err
	}

	var carrierMetrics []*dto.CarrierMetricResponse
	for _, c := range metric.CarrierMetrics {
		carrierMetrics = append(carrierMetrics, &dto.CarrierMetricResponse{
			Name:           c.Name,
			Total:          c.Total,
			FinalPriceMean: c.FinalPriceMean,
			FinalPriceSum:  c.FinalPriceSum,
		})
	}

	return &dto.MetricResponse{
		CheaperShipping:       metric.CheaperShipping,
		MostExpensiveShipping: metric.MostExpensiveShipping,
		CarrierMetrics:        carrierMetrics,
	}, nil
}
