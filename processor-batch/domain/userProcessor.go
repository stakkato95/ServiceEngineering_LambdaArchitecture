package domain

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stakkato95/lambda-architecture/processor-batch/config"
	"github.com/stakkato95/lambda-architecture/processor-batch/logger"
)

type UserProcessor interface {
	DoProcessing() error
}

const partition = 0
const msgBufferSize = 10e3 //10KB

type mysqlUserProcessor struct {
	db *sqlx.DB
}

func NewUserProcessor() UserProcessor {
	logger.Info("connection string: " + config.AppConfig.DbSource)
	db, err := sqlx.Open(config.AppConfig.DbDriver, config.AppConfig.DbSource)
	if err != nil {
		logger.Fatal("can not connect to database: " + err.Error())
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &mysqlUserProcessor{db}
}

func (p *mysqlUserProcessor) DoProcessing() error {
	users, err := p.readAll()
	if err != nil {
		return err
	}

	if len(users) == 0 {
		logger.Info("no users found")
		return nil
	}

	longestNameUser := users[0]
	for _, user := range users {
		if len(user.Name) > len(longestNameUser.Name) {
			longestNameUser = user
		}
	}

	if err = p.cleanTable(); err != nil {
		return err
	}

	if err = p.writeUser(longestNameUser); err != nil {
		return err
	}

	return nil
}

func (p *mysqlUserProcessor) readAll() ([]User, error) {
	users := []User{}
	findAllSql := "SELECT id, name FROM user"
	if err := p.db.Select(&users, findAllSql); err != nil {
		return nil, err
	}
	return users, nil
}

func (p *mysqlUserProcessor) cleanTable() error {
	truncateTable := "TRUNCATE TABLE user_longest_name"
	if _, err := p.db.Exec(truncateTable); err != nil {
		return err
	}
	return nil
}

func (p *mysqlUserProcessor) writeUser(user User) error {
	insertUser := "INSERT INTO user_longest_name (id, name) VALUES (?, ?)"
	result := p.db.MustExec(insertUser, user.Id, user.Name)
	if rows, err := result.RowsAffected(); err != nil {
		return errors.New(fmt.Sprintf("error when inserting user_longest_name: %v", err))
	} else if rows == 0 {
		return errors.New(fmt.Sprintf("no rows affected when inserting user_longest_name: %v", user))
	}

	logger.Info(fmt.Sprintf("updatet user_longest_name: %v", user))
	return nil
}
