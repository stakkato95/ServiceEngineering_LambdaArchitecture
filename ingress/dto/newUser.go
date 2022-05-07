package dto

import "github.com/stakkato95/lambda-architecture/ingress/domain"

type NewUser struct {
	Id   string //`json:"id"`
	Name string //`json:"name"`
}

func (u *NewUser) ToEntity() domain.User {
	return domain.User{
		Id:   u.Id,
		Name: u.Name,
	}
}
