package service

import (
	"github.com/stakkato95/lambda-architecture/ingress/domain"
	"github.com/stakkato95/lambda-architecture/ingress/errs"
)

type userServiceStub struct {
}

func (s *userServiceStub) InjestUser(user domain.User) *errs.AppError {
	return nil
}

func NewUserServiceStub() UserService {
	return &userServiceStub{}
}
