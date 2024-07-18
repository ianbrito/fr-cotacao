package service

import (
	"context"
	"github.com/ianbrito/fr-cotacao/internal/adapter"
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	"github.com/ianbrito/fr-cotacao/internal/dto"
	"github.com/ianbrito/fr-cotacao/internal/infra/repository"
	fr "github.com/ianbrito/fr-cotacao/pkg/frete-rapido"
)

type QuoteService struct {
	ctx        context.Context
	Repository *repository.SQLDispatcherRepository
	Adapter    *adapter.FreteRapidoQuoteAdapter
}

func NewQuoteService(ctx context.Context) *QuoteService {
	return &QuoteService{
		ctx:        ctx,
		Repository: repository.NewSQLDispatcherRepository(ctx, nil),
	}
}

func (qs *QuoteService) GetQuoteService(request *dto.QuoteRequest) (*dto.QuoteResponse, error) {
	client := fr.NewFRCotacaoClient()
	quoteRequest := dto.NewFreteRapidoQuoteRequest(request)

	quotes, err := client.Simulate(quoteRequest.MakeRequest())
	if err != nil {
		return nil, err
	}

	qs.Adapter = adapter.NewFreteRapidoQuoteAdapter(quotes)
	dispatchers, err := qs.SaveQuotes()
	if err != nil {
		return nil, err
	}
	response := dto.NewQuoteResponse(dispatchers)
	return response, nil
}

func (qs *QuoteService) SaveQuotes() ([]*entity.Dispatcher, error) {
	var dispatchers []*entity.Dispatcher
	for _, d := range qs.Adapter.ToEntity() {
		dispatcher, err := qs.Repository.Save(d)
		if err != nil {
			return nil, err
		}
		dispatchers = append(dispatchers, dispatcher)
	}
	return dispatchers, nil
}
