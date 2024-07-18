package dto

import (
	"fmt"
	fr "github.com/ianbrito/fr-cotacao/pkg/frete-rapido"
	"log"
	"os"
	"strconv"
)

type FreteRapidoQuoteRequest struct {
	QuoteResquest *QuoteRequest
}

func NewFreteRapidoQuoteRequest(qr *QuoteRequest) *FreteRapidoQuoteRequest {
	return &FreteRapidoQuoteRequest{
		QuoteResquest: qr,
	}
}

func (r *FreteRapidoQuoteRequest) mapVolumes() ([]*fr.Volume, error) {
	var result []*fr.Volume
	for _, v := range r.QuoteResquest.Volumes {
		result = append(result, &fr.Volume{
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

func (r *FreteRapidoQuoteRequest) createDispatcher(volumes []*fr.Volume) fr.Dispatcher {

	dispatcherZipCode, err := parseZipCode(os.Getenv("TEST_ZIP_CODE"))
	if err != nil {
		log.Printf("Error parsing zip code: %s\n", err)
	}

	return fr.Dispatcher{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		ZipCode:          dispatcherZipCode,
		Volumes:          volumes,
	}
}

func (r *FreteRapidoQuoteRequest) createShipper() fr.Shipper {
	return fr.Shipper{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		Token:            os.Getenv("TEST_TOKEN"),
		PlatformCode:     os.Getenv("TEST_PLATFORM_CODE"),
	}
}

func (r *FreteRapidoQuoteRequest) createRecipient() fr.Recipient {

	recipientZipCode, err := parseZipCode(r.QuoteResquest.Recipient.Address.ZipCode)
	if err != nil {
		log.Printf("Error parsing zip code: %s\n", err)
	}

	return fr.Recipient{
		RegisteredNumber: os.Getenv("TEST_CPNJ"),
		Country:          "BRA",
		ZipCode:          recipientZipCode,
	}
}

func (r *FreteRapidoQuoteRequest) MakeRequest() *fr.QuoteRequest {
	volumes, _ := r.mapVolumes()
	req := &fr.QuoteRequest{
		Shipper:        r.createShipper(),
		Recipient:      r.createRecipient(),
		SimulationType: []int{0},
	}
	req.AddDispatcher(r.createDispatcher(volumes))
	return req
}

func parseZipCode(zipCode string) (uint32, error) {
	parsedZipCode, err := strconv.ParseUint(zipCode, 10, 32)
	if err != nil {
		fmt.Printf("Error parsing zip code: %s\n", err)
		return 0, err
	}
	return uint32(parsedZipCode), nil
}
