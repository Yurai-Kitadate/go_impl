package repository

import (
	"../model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUser(id int64) (*model.User, error)
}

type userRepositoryImpl struct {
	db *sqlx.DB
}
