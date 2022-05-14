package app

import (
	"encoding/json"
	"net/http"

	"github.com/stakkato95/lambda-architecture/service-analytics/dto"
	"github.com/stakkato95/lambda-architecture/service-analytics/logger"
	"github.com/stakkato95/lambda-architecture/service-analytics/service"
)

type AnalyticsHandlers struct {
	service service.AnalyticsService
}

func (h *AnalyticsHandlers) GetUserAnalytics(w http.ResponseWriter, r *http.Request) {
	user, counts, err := h.service.GetAnalytics()
	if err != nil {
		logger.Error("error when fetching analytics: " + err.Error())
		writeResponse(w, http.StatusInternalServerError, "error when fetching analytics")
	}

	writeResponse(w, http.StatusOK, dto.ToDto(user, counts))
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	//1 Content-Type - always first
	w.Header().Add("Content-Type", "application/json")
	//2 HTTP status code
	w.WriteHeader(code)
	//3 body
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Fatal("can not encode data: " + err.Error())
	}
}
