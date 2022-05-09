package service

import (
	"github.com/stakkato95/lambda-architecture/ingress/domain"
	"github.com/stakkato95/lambda-architecture/ingress/errs"
)

type UserService interface {
	InjestUser(domain.User) *errs.AppError
}

type simpleUserService struct {
	repo domain.UserRepository
}

func (s *simpleUserService) InjestUser(user domain.User) *errs.AppError {
	return s.repo.InjestUser(user)
}

func NewSimpleUserService(repo domain.UserRepository) UserService {
	return &simpleUserService{repo}
}
