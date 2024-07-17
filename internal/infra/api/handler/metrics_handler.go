package handler

import (
	"context"
	"fmt"
	"github.com/go-chi/render"
	"github.com/ianbrito/fr-cotacao/internal/dto"
	"github.com/ianbrito/fr-cotacao/internal/service"
	"net/http"
	"strconv"
)

type MetricsHandler struct {
	Service *service.MetricService
}

func NewMetricsHandler(ctx context.Context) *MetricsHandler {
	return &MetricsHandler{
		Service: service.NewMetricService(ctx),
	}
}

func (h *MetricsHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Query().Has("last_quotes") {
		limit := r.URL.Query().Get("last_quotes")

		l, err := strconv.ParseInt(limit, 10, 32)
		if err != nil {
			fmt.Print(err)
			response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusBadRequest)
			render.Render(w, r, response)
			return
		}

		response, err := h.Service.GetMetricsWithLimit(int(l))
		if err != nil {
			response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusInternalServerError)
			render.Render(w, r, response)
			return
		}

		render.Render(w, r, response)
		return
	}

	response, err := h.Service.GetMetrics()
	if err != nil {
		response := dto.NewErrorResponse("Ocorreu um erro ao processar sua solicitação!", http.StatusInternalServerError)
		render.Render(w, r, response)
		return
	}

	render.Render(w, r, response)
}
