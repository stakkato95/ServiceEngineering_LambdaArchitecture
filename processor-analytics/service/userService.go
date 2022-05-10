package service

import "github.com/stakkato95/lambda-architecture/processor-analytics/domain"

// "github.com/stakkato95/lambda-architecture/ingress/domain"
// "github.com/stakkato95/lambda-architecture/ingress/errs"

type UserService interface {
	GetUserCount() int
}

type simpleUserService struct {
	repo domain.UserProcessor
}

func (s *simpleUserService) GetUserCount() int {
	return s.repo.GetUserCount()
}

func NewUserService(repo domain.UserProcessor) UserService {
	return &simpleUserService{repo}
}
