package dto

import "github.com/stakkato95/lambda-architecture/service-analytics/domain"

type UserCountDto struct {
	Id        string `json:"id"`
	Time      string `json:"time"`
	UserCount int    `json:"userCount"`
}

func UserCountToDto(userCount *domain.UserCount) UserCountDto {
	return UserCountDto{
		Id:        userCount.Id,
		Time:      userCount.Time,
		UserCount: userCount.UserCount,
	}
}
