package errs

import "net/http"

type AppError struct {
	Code int    `json:"omitempty"`
	Msg  string `json:",msg"`
}

func NewInjestError(msg string) *AppError {
	return &AppError{
		http.StatusInternalServerError,
		msg,
	}
}
