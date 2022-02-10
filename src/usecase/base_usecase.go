package usecase

import (
	"go-api/src/helpers/bbl"
	bblModel "go-api/src/helpers/bbl/model"
	"go-api/src/helpers/shape"
	"go-api/src/model"
	"go-api/src/repository"
)

type Usecase struct {
	repos             repository.Interface
	helperShapeRec    shape.Interface
	helperShapeCircle shape.Interface
	bbl               bbl.Service
}

func New(rp repository.Interface, hpR, hpC shape.Interface, bbl bbl.Service) *Usecase {
	return &Usecase{
		repos:             rp,
		helperShapeRec:    hpR,
		helperShapeCircle: hpC,
		bbl:               bbl,
	}
}

type Interface interface {
	GetUser(id int32) (output model.Users, err error)
	GetUsers(filter model.Filter) (output []model.Users, paginate model.Paginate, err error)
	AddUser(input model.Users) (output model.UserOutput, err error)
	UpdateUser(input model.Users, id int32) (output model.Users, err error)
	DeleteUser(id int32) error

	RegisterBBLCard(input model.RegisterCard) (output model.RegisterCard, err error)
	CallbackDataFeed(input model.DataFeed) (output model.DataFeed, err error)
	GetMemberPay(input bblModel.BBLRequest) (output []bblModel.MemberPayAccounts, err error)
}
