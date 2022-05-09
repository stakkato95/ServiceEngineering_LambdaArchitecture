package service

import "github.com/stakkato95/lambda-architecture/processor-analytics/domain"

// "github.com/stakkato95/lambda-architecture/ingress/domain"
// "github.com/stakkato95/lambda-architecture/ingress/errs"

type UserService interface {
	GetUserCount() int
}

type simpleUserService struct {
	repo domain.UserRepository
}

func (s *simpleUserService) GetUserCount() int {
	return s.repo.GetUserCount()
}

func NewUserService(repo domain.UserRepository) UserService {
	return &simpleUserService{repo}
}
