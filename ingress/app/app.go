package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stakkato95/lambda-architecture/ingress/config"
	"github.com/stakkato95/lambda-architecture/ingress/domain"
	"github.com/stakkato95/lambda-architecture/ingress/logger"
	"github.com/stakkato95/lambda-architecture/ingress/service"
)

func Start() {
	router := mux.NewRouter()

	repo := domain.NewKafkaUserRepository()
	defer func() {
		repo.Destroy()
		logger.Info("Kafka connection successfully destroyed")
	}()

	service := service.NewSimpleUserService(repo)
	handlers := UserHandlers{service}

	router.HandleFunc("/user", handlers.CreateUser).Methods(http.MethodPost)

	// u := domain.User{Id: "id", Name: "myname"}
	// repo.InjestUser(u)

	port := config.AppConfig.ServerPort
	logger.Info(fmt.Sprintf("started ingress at %s", port))
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Fatal("error when starting ingress " + err.Error())
	}
}
