package frete_rapido

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ianbrito/fr-cotacao/pkg/frete-rapido/models"
	"log"
	"net/http"
	"os"
)

type FRCotacaoClient struct {
	Endpoint string
}

func NewFRCotacaoClient() *FRCotacaoClient {
	return &FRCotacaoClient{
		Endpoint: os.Getenv("FR_COTACAO_ENDPOINT"),
	}
}

func (c *FRCotacaoClient) Simulate(request *models.QuoteRequest) (*models.Response, error) {
	response, err := http.Post(c.Endpoint, "application/json", bytes.NewReader(request.ParseToJSON()))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}

	var data models.Response
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return &data, nil
}
