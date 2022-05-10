package domain

import (
	"fmt"

	"github.com/stakkato95/lambda-architecture/service-transactional/logger"
)

type UserSink interface {
	Sink(User) error
}

type mySqlUserSink struct {
}

func NewUserSink() UserSink {
	return &mySqlUserSink{}
}

func (s *mySqlUserSink) Sink(user User) error {
	logger.Info(fmt.Sprintf("passed to sink: %v", user))
	return nil
}
