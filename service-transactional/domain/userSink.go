package domain

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stakkato95/lambda-architecture/service-transactional/config"
	"github.com/stakkato95/lambda-architecture/service-transactional/logger"
)

type UserSink interface {
	Sink(User) error
}

type mySqlUserSink struct {
	db *sqlx.DB
}

func NewUserSink() UserSink {
	logger.Info("connection string: " + config.AppConfig.DbSource)
	db, err := sqlx.Open(config.AppConfig.DbDriver, config.AppConfig.DbSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &mySqlUserSink{db}
}

func (s *mySqlUserSink) Sink(user User) error {
	insertUser := "INSERT INTO user (id, name) VALUES (?, ?)"
	result := s.db.MustExec(insertUser, user.Id, user.Name)
	if rows, err := result.RowsAffected(); err != nil {
		return errors.New(fmt.Sprintf("error when inserting user: %v", err))
	} else if rows == 0 {
		return errors.New(fmt.Sprintf("no rows affected when inserting %v", user))
	}

	logger.Info(fmt.Sprintf("passed to sink: %v", user))
	return nil
}
