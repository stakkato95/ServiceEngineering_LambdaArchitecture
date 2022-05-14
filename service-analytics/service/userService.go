package service

import "github.com/stakkato95/lambda-architecture/service-analytics/domain"

type UserService interface {
	GetLongestNameUser() (domain.User, error)
}

type simpleUserService struct {
}

func NewUserService() UserService {
	return &simpleUserService{}
}

func (s *simpleUserService) GetLongestNameUser() (domain.User, error) {
	return domain.User{Id: "1", Name: "hello"}, nil
}
