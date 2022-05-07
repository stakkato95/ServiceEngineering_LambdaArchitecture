package app

import (
	"encoding/json"
	"net/http"

	"github.com/stakkato95/lambda-architecture/ingress/service"
)

type UserHandlers struct {
	service service.UserService
}

func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "ingressWorks")
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
