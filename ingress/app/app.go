package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stakkato95/lambda-architecture/ingress/config"
	"github.com/stakkato95/lambda-architecture/ingress/logger"
	"github.com/stakkato95/lambda-architecture/ingress/service"
)

func Start() {
	router := mux.NewRouter()

	handlers := UserHandlers{service.NewUserServiceStub()}

	router.HandleFunc("/user", handlers.CreateUser).Methods(http.MethodPost)

	port := config.AppConfig.ServerPort
	logger.Info(fmt.Sprintf("started ingress at %s", port))
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Error("error when starting ingress " + err.Error())
	}
}
