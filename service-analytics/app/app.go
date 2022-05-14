package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stakkato95/lambda-architecture/service-analytics/config"
	"github.com/stakkato95/lambda-architecture/service-analytics/domain"
	"github.com/stakkato95/lambda-architecture/service-analytics/logger"
	"github.com/stakkato95/lambda-architecture/service-analytics/service"
)

func Start() {
	router := mux.NewRouter()

	userRepo := domain.NewUserRepository()
	userCountRepo := domain.NewUserCountRepository()

	userService := service.NewUserService(userRepo)
	userCountService := service.NewUserCountService(userCountRepo)
	service := service.NewAnalyticsService(userService, userCountService)

	handlers := AnalyticsHandlers{service}

	router.HandleFunc("/analytics", handlers.GetUserAnalytics).Methods(http.MethodGet)

	port := config.AppConfig.ServerPort
	logger.Info(fmt.Sprintf("started service-analytics at %s", port))
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Fatal("error when starting ingress " + err.Error())
	}
}
