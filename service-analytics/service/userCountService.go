package service

import "github.com/stakkato95/lambda-architecture/service-analytics/domain"

type UserCountService interface {
	GetUserCount() ([]domain.UserCount, error)
}

type simpleUserCountService struct {
	repo domain.UserCountRepository
}

func NewUserCountService(repo domain.UserCountRepository) UserCountService {
	return &simpleUserCountService{repo}
}

func (s *simpleUserCountService) GetUserCount() ([]domain.UserCount, error) {
	return s.repo.GetUserCounts(), nil
}
