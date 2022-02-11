package repository

import (
	"github.com/kraingkrai-k/go-api/src/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type Interface interface {
	// user_usecase.go
	GetUser(id int32) (output model.Users, err error)
	GetUsers(filter model.Filter) (output []model.Users, paginate model.Paginate, err error)
	AddUser(input model.Users) (output model.UserOutput, err error)
	UpdateUser(input model.Users, id int32) (output model.Users, err error)
	DeleteUser(id int32) error
}
