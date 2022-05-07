package service

import (
	"github.com/stakkato95/lambda-architecture/ingress/domain"
	"github.com/stakkato95/lambda-architecture/ingress/errs"
	"github.com/stakkato95/lambda-architecture/ingress/logger"
)

type userServiceStub struct {
}

func (s *userServiceStub) InjestUser(user domain.User) *errs.AppError {
	logger.Info("stub injest")
	return nil
}

func NewUserServiceStub() UserService {
	return &userServiceStub{}
}
