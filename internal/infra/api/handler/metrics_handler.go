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

// GetMetrics
// @Summary Get metrics
// @Description Retrieve metrics, optionally limited by the number of last quotes
// @Tags metrics
// @Accept json
// @Produce json
// @Param last_quotes query int false "Number of last quotes to limit the metrics"
// @Success 200 {object} dto.MetricResponse
// @Failure 400 {object} dto.ErrorResponse "Invalid request"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /metrics [get]
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
