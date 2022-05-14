package dto

import "github.com/stakkato95/lambda-architecture/service-analytics/domain"

type UserDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func UserToDto(user *domain.User) UserDto {
	return UserDto{
		Id:   user.Id,
		Name: user.Name,
	}
}
