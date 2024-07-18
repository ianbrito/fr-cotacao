package frete_rapido

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *FRCotacaoClient) Simulate(request *QuoteRequest) (*QuoteResponse, error) {

	response, err := http.Post(c.Endpoint, "application/json", bytes.NewReader(request.ParseToJSON()))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}

	var data QuoteResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
