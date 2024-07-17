package repository

import (
	"context"
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	"github.com/ianbrito/fr-cotacao/internal/infra/db"
)

type MetricRepository interface {
	GetMetric() (*entity.Metric, error)
	GetCarrierMetrics() ([]*entity.CarrierMetric, error)
	GetMetricWithLimit(limit int) (*entity.Metric, error)
	GetCarrierMetricsWithLimit(ids []string) ([]*entity.CarrierMetric, error)
}

type SQLMetricRepository struct {
	Queries *db.Queries
	Ctx     context.Context
}

func NewSQLMetricRepository(ctx context.Context, queries *db.Queries) *SQLMetricRepository {
	if queries == nil {
		conn := db.GetConnection()
		queries = db.New(conn)
	}
	return &SQLMetricRepository{
		Queries: queries,
		Ctx:     ctx,
	}
}

func (r *SQLMetricRepository) GetCarrierMetrics() ([]*entity.CarrierMetric, error) {
	metrics, err := r.Queries.GetCarrierMetric(r.Ctx)
	if err != nil {
		return nil, err
	}
	var carrierMetrics []*entity.CarrierMetric
	for _, m := range metrics {
		carrierMetrics = append(carrierMetrics, &entity.CarrierMetric{
			Name:           m.CarrierName,
			Total:          m.Total,
			FinalPriceMean: m.FinalPriceMean,
			FinalPriceSum:  m.FinalPriceSum,
		})
	}
	return carrierMetrics, nil
}

func (r *SQLMetricRepository) GetMetric() (*entity.Metric, error) {
	m, err := r.Queries.GetPriceMetric(r.Ctx)
	if err != nil {
		return nil, err
	}

	carrierMetrics, err := r.GetCarrierMetrics()
	if err != nil {
		return nil, err
	}

	metric := &entity.Metric{
		CheaperShipping:       m.CheaperShipping,
		MostExpensiveShipping: m.MostExpensiveShipping,
		CarrierMetrics:        carrierMetrics,
	}

	return metric, nil
}

func (r *SQLMetricRepository) GetCarrierMetricsWithLimit(ids []string) ([]*entity.CarrierMetric, error) {
	metrics, err := r.Queries.GetCarrierMetricsWithLimit(r.Ctx, ids)
	if err != nil {
		return nil, err
	}
	var carrierMetrics []*entity.CarrierMetric
	for _, m := range metrics {
		carrierMetrics = append(carrierMetrics, &entity.CarrierMetric{
			Name:           m.CarrierName,
			Total:          m.Total,
			FinalPriceMean: m.FinalPriceMean,
			FinalPriceSum:  m.FinalPriceSum,
		})
	}
	return carrierMetrics, nil
}

func (r *SQLMetricRepository) GetMetricWithLimit(limit int) (*entity.Metric, error) {

	ids, err := r.Queries.GetDispatcherIdsWithLimit(r.Ctx, int32(limit))
	if err != nil {
		return nil, err
	}

	m, err := r.Queries.GetPriceMetricWithLimit(r.Ctx, ids)
	if err != nil {
		return nil, err
	}

	carrierMetrics, err := r.GetCarrierMetricsWithLimit(ids)
	if err != nil {
		return nil, err
	}

	metric := &entity.Metric{
		CheaperShipping:       m.CheaperShipping,
		MostExpensiveShipping: m.MostExpensiveShipping,
		CarrierMetrics:        carrierMetrics,
	}

	return metric, nil
}
