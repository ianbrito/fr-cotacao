package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type AddressRequest struct {
	ZipCode string `json:"zipcode" validate:"required"`
}

type RecipientRequest struct {
	Address *AddressRequest `json:"address" validate:"required"`
}

type VolumeRequest struct {
	Category      int     `json:"category" validate:"required"`
	Amount        int     `json:"amount" validate:"required"`
	UnitaryWeight float64 `json:"unitary_weight" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	Sku           string  `json:"sku" validate:"required"`
	Height        float64 `json:"height" validate:"required"`
	Width         float64 `json:"width" validate:"required"`
	Length        float64 `json:"length" validate:"required"`
}

type QuoteRequest struct {
	Recipient *RecipientRequest `json:"recipient" validate:"required"`
	Volumes   []*VolumeRequest  `json:"volumes" validate:"required,dive"`
}

func validateRequest(v interface{}) []string {
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := strings.TrimPrefix(err.StructNamespace(), "QuoteRequest.")
			errors = append(errors, fmt.Sprintf("%s: %s", fieldName, err.Tag()))
		}
		return errors
	}
	return nil
}

func Quote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var quote QuoteRequest

	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if errors := validateRequest(&quote); errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": errors,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}
