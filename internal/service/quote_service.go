package service

import (
	"context"
	"fmt"
	"github.com/ianbrito/fr-cotacao/internal/domain/entity"
	"github.com/ianbrito/fr-cotacao/internal/dto"
	"github.com/ianbrito/fr-cotacao/internal/infra/repository"
	frete_rapido "github.com/ianbrito/fr-cotacao/pkg/frete-rapido"
	"github.com/ianbrito/fr-cotacao/pkg/frete-rapido/models"
	"log"
	"os"
	"strconv"
)

type QuoteService struct {
	ctx        context.Context
	Repository *repository.SQLDispatcherRepository
}

func NewQuoteService(ctx context.Context) *QuoteService {
	return &QuoteService{
		ctx:        ctx,
		Repository: repository.NewSQLDispatcherRepository(ctx, nil),
	}
}

func (qs *QuoteService) GetQuoteService(request *dto.QuoteRequest) (*dto.QuoteResponse, error) {
	client := frete_rapido.NewFRCotacaoClient()

	volumes, err := mapVolumes(request.Volumes)
	if err != nil {
		return nil, err
	}

	quote := createQuoteRequest(request)
	quote.AddDispatcher(createDispatcher(volumes))

	quotes, err := client.Simulate(quote)
	if err != nil {
		return nil, err
	}

	qs.SaveQuotes(quotes)

	response := createQuoteResponse(quotes)
	return response, nil
}

func (qs *QuoteService) SaveQuotes(quotes *models.Response) error {
	for _, d := range quotes.Dispatchers {
		var offers []*entity.Offer
		for _, o := range d.Offers {
			carrier := &entity.Carrier{
				Reference:        int64(o.Carrier.Reference),
				Name:             o.Carrier.Name,
				RegisteredNumber: o.Carrier.RegisteredNumber,
				StateInscription: o.Carrier.StateInscription,
				LogoUrl:          o.Carrier.Logo,
			}
			weights := &entity.Weights{
				Real:  o.Weights.Real,
				Cubed: &o.Weights.Cubed,
				Used:  o.Weights.Cubed,
			}

			deliveryTime := &entity.DeliveryTime{
				Days:          o.DeliveryTime.Days,
				Hours:         o.DeliveryTime.Hours,
				Minutes:       o.DeliveryTime.Days,
				EstimatedDate: o.DeliveryTime.EstimatedDate,
			}
			originalDeliveryTime := &entity.DeliveryTime{
				Days:          o.OriginalDeliveryTime.Days,
				Hours:         o.OriginalDeliveryTime.Hours,
				Minutes:       o.OriginalDeliveryTime.Days,
				EstimatedDate: o.OriginalDeliveryTime.EstimatedDate,
			}

			offers = append(offers, entity.NewOffer(
				int32(o.Offer),
				int32(o.SimulationType),
				carrier,
				o.Service,
				o.ServiceCode,
				o.ServiceDescription,
				o.Identifier,
				o.DeliveryNote,
				o.HomeDelivery,
				o.CarrierNeedsToReturnToSender,
				o.Expiration,
				o.CostPrice,
				o.FinalPrice,
				o.Modal,
				weights,
				deliveryTime,
				originalDeliveryTime,
			))
		}
		_, err := qs.Repository.Save(&entity.Dispatcher{
			ID:                         d.ID,
			RequestID:                  d.RequestID,
			RegisteredNumberShipper:    d.RegisteredNumberShipper,
			RegisteredNumberDispatcher: d.RegisteredNumberDispatcher,
			ZipCode:                    d.ZipcodeOrigin,
			Offers:                     offers,
		})

		if err != nil {
			return err
		}
	}
	return nil
}

func mapVolumes(volumes []*dto.VolumeRequest) ([]*models.Volume, error) {
	var result []*models.Volume
	for _, v := range volumes {
		result = append(result, &models.Volume{
			Category:      strconv.Itoa(v.Category),
			Amount:        v.Amount,
			UnitaryWeight: v.UnitaryWeight,
			UnitaryPrice:  v.Price,
			Sku:           v.Sku,
			Height:        v.Height,
			Width:         v.Width,
			Length:        v.Length,
		})
	}
	return result, nil
}

func parseZipCode(zipCode string) (uint32, error) {
	parsedZipCode, err := strconv.ParseUint(zipCode, 10, 32)
	if err != nil {
		fmt.Printf("Error parsing zip code: %s\n", err)
		return 0, err
	}
	return uint32(parsedZipCode), nil
}

func createShipper() models.Shipper {
	return models.Shipper{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		Token:            os.Getenv("TEST_TOKEN"),
		PlatformCode:     os.Getenv("TEST_PLATFORM_CODE"),
	}
}

func createRecipient(qr *dto.QuoteRequest) models.Recipient {

	recipientZipCode, err := parseZipCode(qr.Recipient.Address.ZipCode)
	if err != nil {
		log.Printf("Error parsing zip code: %s\n", err)
	}

	return models.Recipient{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		Country:          "BRA",
		ZipCode:          recipientZipCode,
	}
}

func createQuoteRequest(qr *dto.QuoteRequest) *models.QuoteRequest {
	return &models.QuoteRequest{
		Shipper:        createShipper(),
		Recipient:      createRecipient(qr),
		SimulationType: []int{0},
	}
}

func createDispatcher(volumes []*models.Volume) models.Dispatcher {

	dispatcherZipCode, err := parseZipCode(os.Getenv("TEST_ZIP_CODE"))
	if err != nil {
		log.Printf("Error parsing zip code: %s\n", err)
	}

	return models.Dispatcher{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		ZipCode:          dispatcherZipCode,
		Volumes:          volumes,
	}
}

func createQuoteResponse(data *models.Response) *dto.QuoteResponse {
	response := dto.NewQuoteResponse()
	for _, dispatcher := range data.Dispatchers {
		for _, offer := range dispatcher.Offers {
			response.AddCarrier(&dto.CarrierResponse{
				Name:     offer.Carrier.Name,
				Service:  offer.Modal,
				Deadline: strconv.Itoa(offer.DeliveryTime.Days),
				Price:    offer.FinalPrice,
			})
		}
	}
	return response
}
