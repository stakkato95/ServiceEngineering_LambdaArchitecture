package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stakkato95/lambda-architecture/ingress/config"
	"github.com/stakkato95/lambda-architecture/ingress/logger"
)

func Start() {
	router := mux.NewRouter()

	handlers := UserHandlers{}

	router.HandleFunc("/user", handlers.CreateUser).Methods(http.MethodPost)

	uri := config.ServerUri()

	logger.Info(fmt.Sprintf("started ingress at %s", uri))
	if err := http.ListenAndServe(uri, router); err != nil {
		logger.Error("error when starting ingress " + err.Error())
	}
}
