package service

import (
	"fmt"
	"github.com/ianbrito/fr-cotacao/internal/dto"
	frete_rapido "github.com/ianbrito/fr-cotacao/pkg/frete-rapido"
	"github.com/ianbrito/fr-cotacao/pkg/frete-rapido/models"
	"strconv"
)

func GetQuoteService(request *dto.QuoteRequest) (*dto.QuoteResponse, error) {
	client := frete_rapido.NewFRCotacaoClient()

	s, r, d := request.ToEntity()

	var volumes []*models.Volume
	for _, v := range d.Volumes {
		volumes = append(volumes, &models.Volume{
			Category:      v.Category,
			Amount:        v.Amount,
			UnitaryWeight: v.UnitaryWeight,
			UnitaryPrice:  v.Price,
			Sku:           v.Sku,
			Height:        v.Height,
			Width:         v.Width,
			Length:        v.Length,
		})
	}

	dispatcherZipCode, err := strconv.ParseUint(d.ZipCode, 10, 32)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	recipientZipCode, err := strconv.ParseUint(r.Address.ZipCode, 10, 32)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	quote := &models.QuoteRequest{
		Shipper: models.Shipper{
			Token:            s.Token,
			RegisteredNumber: s.RegisteredNumber,
			PlatformCode:     s.PlatformCode,
		},
		Recipient: models.Recipient{
			RegisteredNumber: r.RegisteredNumber,
			Country:          r.Address.Country,
			ZipCode:          uint32(recipientZipCode),
		},
		SimulationType: []int{0},
	}
	quote.AddDispatcher(models.Dispatcher{
		RegisteredNumber: d.RegisteredNumber,
		ZipCode:          uint32(dispatcherZipCode),
		Volumes:          volumes,
	})

	var data *models.Response
	data, err = client.Simulate(quote)
	if err != nil {
		return nil, err
	}

	response := dto.NewQuoteResponse()
	for _, dispatcher := range data.Dispatchers {
		for _, o := range dispatcher.Offers {
			response.AddCarrier(&dto.CarrierResponse{
				Name:     o.Carrier.Name,
				Service:  o.Modal,
				Deadline: "1",
				Price:    o.FinalPrice,
			})
		}
	}
	return response, nil
}
