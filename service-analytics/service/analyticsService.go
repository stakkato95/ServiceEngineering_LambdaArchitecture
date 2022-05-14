package service

import "github.com/stakkato95/lambda-architecture/service-analytics/domain"

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
	if err != nil {
		return nil, nil, err
	}

	return &user, counts, nil
}
