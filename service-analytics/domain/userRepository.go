package domain

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stakkato95/lambda-architecture/service-analytics/config"
	"github.com/stakkato95/lambda-architecture/service-analytics/logger"
)

type UserRepository interface {
	GetLongestNameUser() (*User, error)
}

type mysqlUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() UserRepository {
	logger.Info("connection string: " + config.AppConfig.DbSource)
	db, err := sqlx.Open(config.AppConfig.DbDriver, config.AppConfig.DbSource)
	if err != nil {
		logger.Fatal("can not connect to database: " + err.Error())
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &mysqlUserRepository{db}
}

func (r *mysqlUserRepository) GetLongestNameUser() (*User, error) {
	users := []User{}
	findAllSql := "SELECT id, name FROM user_longest_name"
	if err := r.db.Select(&users, findAllSql); err != nil {
		return nil, err
	}
	if len(users) != 1 {
		return nil, errors.New("more than one user with longest name")
	}

	return &users[0], nil
}
