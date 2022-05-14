package domain

import "time"

type UserCount struct {
	Id        string
	Time      time.Time
	UserCount int
}
