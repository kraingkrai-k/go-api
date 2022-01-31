package usecase

import (
	"go-api/src/model"
	"go-api/src/repository"
)

type Usecase struct {
	repos repository.Interface
}

func New(rp repository.Interface) *Usecase {
	return &Usecase{
		repos: rp,
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
