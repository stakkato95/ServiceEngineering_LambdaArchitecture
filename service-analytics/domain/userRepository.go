package domain

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetLongestNameUser() (*User, error)
}

type mysqlUserRepository struct {
	db *sqlx.DB
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
