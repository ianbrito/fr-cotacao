package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ianbrito/fr-cotacao/internal/dto"
	"github.com/ianbrito/fr-cotacao/internal/service"
	"net/http"
	"strings"
)

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

func GetQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var quote dto.QuoteRequest
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

	err := service.GetQuoteService(&quote)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
