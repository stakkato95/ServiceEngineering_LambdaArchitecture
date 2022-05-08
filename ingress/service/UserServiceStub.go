package service

import (
	"github.com/stakkato95/lambda-architecture/ingress/domain"
	"github.com/stakkato95/lambda-architecture/ingress/errs"
)

type userServiceStub struct {
	repo domain.UserRepository
}

func (s *userServiceStub) InjestUser(user domain.User) *errs.AppError {
	return s.repo.InjestUser(user)
}

func NewUserServiceStub(repo domain.UserRepository) UserService {
	return &userServiceStub{repo}
}
