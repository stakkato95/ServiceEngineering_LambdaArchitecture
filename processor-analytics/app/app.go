package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stakkato95/lambda-architecture/processor-analytics/config"
	"github.com/stakkato95/lambda-architecture/processor-analytics/logger"
	"github.com/stakkato95/lambda-architecture/processor-analytics/service"
)

func Start() {
	router := mux.NewRouter()

	// repo := domain.NewKafkaUserRepository()
	// defer func() {
	// 	repo.Destroy()
	// 	logger.Info("Kafka connection successfully destroyed")
	// }()

	service := service.NewSimpleUserService( /*repo*/ )
	handlers := UserHandlers{service}

	router.HandleFunc("/user/count", handlers.ReadUserCount).Methods(http.MethodGet)

	port := config.AppConfig.ServerPort
	logger.Info(fmt.Sprintf("started processor-analytics at %s", port))
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Fatal("error when starting ingress " + err.Error())
	}
}
