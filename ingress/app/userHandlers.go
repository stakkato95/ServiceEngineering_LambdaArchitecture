package app

import (
	"encoding/json"
	"net/http"

	"github.com/stakkato95/lambda-architecture/ingress/dto"
	"github.com/stakkato95/lambda-architecture/ingress/logger"
	"github.com/stakkato95/lambda-architecture/ingress/service"
)

type UserHandlers struct {
	service service.UserService
}

func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser dto.NewUser
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		logger.Error("empty body: " + err.Error())
		writeResponse(w, http.StatusBadRequest, "empty body")
		return
	}

	if err := h.service.InjestUser(newUser.ToEntity()); err != nil {
		logger.Error("user injection error " + err.Msg)
		writeResponse(w, err.Code, err.Msg)
	} else {
		writeResponse(w, http.StatusCreated, "user injested")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	//1 Content-Type - always first
	w.Header().Add("Content-Type", "application/json")
	//2 HTTP status code
	w.WriteHeader(code)
	//3 body
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
