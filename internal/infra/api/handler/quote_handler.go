package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/render"
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
		response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusBadRequest)
		render.Render(w, r, response)
		return
	}

	if errors := validateRequest(&quote); errors != nil {
		response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusBadRequest)
		response.Errors = errors
		render.Render(w, r, response)
		return
	}

	response, err := service.GetQuoteService(&quote)
	if err != nil {
		response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusInternalServerError)
		render.Render(w, r, response)
		return
	}

	render.Render(w, r, response)
}
