package handler

import (
	"context"
	"github.com/go-chi/render"
	"github.com/ianbrito/fr-cotacao/internal/dto"
	"github.com/ianbrito/fr-cotacao/internal/service"
	"github.com/ianbrito/fr-cotacao/utils/validator"
	"net/http"
)

type QuoteHandler struct {
	Service *service.QuoteService
}

func NewQuoteHandler(ctx context.Context) *QuoteHandler {
	return &QuoteHandler{
		Service: service.NewQuoteService(ctx),
	}
}

// GetQuote
// @Summary Get a quote
// @Description Get a quote based on the provided request
// @Tags quotes
// @Accept json
// @Produce json
// @Param QuoteRequest body dto.QuoteRequest true "Quote Request"
// @Success 200 {object} dto.QuoteResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /quote [post]
func (qh *QuoteHandler) GetQuote(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var quote dto.QuoteRequest

	if errors := validator.Validate(r.Body, &quote); errors != nil {
		response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusBadRequest)
		response.Errors = errors
		render.Render(w, r, response)
		return
	}

	response, err := qh.Service.GetQuoteService(&quote)
	if err != nil {
		response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusInternalServerError)
		render.Render(w, r, response)
		return
	}

	render.Render(w, r, response)
}
