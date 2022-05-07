package service

import (
	"github.com/stakkato95/lambda-architecture/ingress/domain"
	"github.com/stakkato95/lambda-architecture/ingress/errs"
)

type UserService interface {
	InjestUser(domain.User) *errs.AppError
}
