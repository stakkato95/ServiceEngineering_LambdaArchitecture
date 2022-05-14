package dto

import (
	"github.com/stakkato95/lambda-architecture/service-analytics/domain"
)

type UserAnalyticsDto struct {
	LongestNameUser *UserDto       `json:"longestNameUser"`
	UserCount       []UserCountDto `json:"userCount"`
}

func ToDto(user *domain.User, userCounts []domain.UserCount) UserAnalyticsDto {
	userCountDtos := make([]UserCountDto, len(userCounts))
	for i, userCount := range userCounts {
		userCountDtos[i] = UserCountToDto(&userCount)
	}

	var userDto *UserDto
	if user != nil {
		userDto = UserToDto(user)
	}

	return UserAnalyticsDto{
		LongestNameUser: userDto,
		UserCount:       userCountDtos,
	}
}
