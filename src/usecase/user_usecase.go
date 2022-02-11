package usecase

import (
	"github.com/kraingkrai-k/go-api/src/model"
)

func (uc *Usecase) GetUsers(filter model.Filter) (output []model.Users, paginate model.Paginate, err error) {

	users, paginate, err := uc.repos.GetUsers(filter)
	if err != nil {
		return nil, paginate, err
	}

	return users, paginate, nil
}

func (uc *Usecase) AddUser(input model.Users) (output model.UserOutput, err error) {

	result, err := uc.repos.AddUser(input)
	if err != nil {
		return model.UserOutput{}, err
	}

	return result, nil
}

func (uc *Usecase) UpdateUser(input model.Users, id int32) (output model.Users, err error) {

	input.ID = id
	result, err := uc.repos.UpdateUser(input, id)
	if err != nil {
		return model.Users{}, err
	}

	return result, nil
}

func (uc *Usecase) DeleteUser(id int32) error {

	err := uc.repos.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) GetUser(id int32) (output model.Users, err error) {

	user, err := uc.repos.GetUser(id)
	if err != nil {
		return model.Users{}, err
	}

	return user, nil
}
