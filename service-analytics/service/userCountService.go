package service

import "github.com/stakkato95/lambda-architecture/service-analytics/domain"

type UserCountService interface {
	GetUserCount() ([]domain.UserCount, error)
}

type simpleUserCountService struct {
}

func NewUserCountService() UserCountService {
	return &simpleUserCountService{}
}

func (s *simpleUserCountService) GetUserCount() ([]domain.UserCount, error) {
	counts := []domain.UserCount{
		{Id: "1", Time: "t1", UserCount: 0},
		{Id: "2", Time: "t2", UserCount: 1},
		{Id: "3", Time: "t3", UserCount: 2},
	}
	return counts, nil
}
