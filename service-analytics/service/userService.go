package service

import "github.com/stakkato95/lambda-architecture/service-analytics/domain"

type UserService interface {
	GetLongestNameUser() (*domain.User, error)
}

type simpleUserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return &simpleUserService{repo}
}

func (s *simpleUserService) GetLongestNameUser() (*domain.User, error) {
	if user, err := s.repo.GetLongestNameUser(); err != nil {
		return nil, err
	} else {
		return user, nil
	}
	// return &domain.User{Id: "1", Name: "hello"}, nil
}
