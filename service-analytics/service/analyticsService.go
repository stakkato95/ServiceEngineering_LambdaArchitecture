package service

import (
	"fmt"

	"github.com/stakkato95/lambda-architecture/service-analytics/domain"
	"github.com/stakkato95/lambda-architecture/service-analytics/logger"
)

type AnalyticsService interface {
	GetAnalytics() (*domain.User, []domain.UserCount, error)
}

type simpleAnalyticsService struct {
	userService      UserService
	userCountService UserCountService
}

func NewAnalyticsService(userService UserService, userCountService UserCountService) AnalyticsService {
	return &simpleAnalyticsService{userService: userService, userCountService: userCountService}
}

func (s *simpleAnalyticsService) GetAnalytics() (*domain.User, []domain.UserCount, error) {
	user, err := s.userService.GetLongestNameUser()
	if err != nil {
		return nil, nil, err
	}

	counts, err := s.userCountService.GetUserCount()
	logger.Info(fmt.Sprintf("counts %d", len(counts)))
	if err != nil {
		return nil, nil, err
	}

	return user, counts, nil
}
