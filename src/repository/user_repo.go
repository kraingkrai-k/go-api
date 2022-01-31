package repository

import (
	"errors"
	"go-api/src/model"

	"gorm.io/gorm"
)

func (repo *Repository) AddUser(input model.Users) (output model.UserOutput, err error) {

	if err = repo.db.Create(&input).Scan(&output).Error; err != nil {
		return model.UserOutput{}, err
	}

	return output, nil
}

func (repo *Repository) GetUsers(filter model.Filter) (output []model.Users, paginate model.Paginate, err error) {

	var count int64
	var userModel model.Users

	result := repo.db.Scopes(userModel.Filter(filter)).Find(&output).Count(&count)
	if result.Error != nil {
		return nil, paginate, err
	}

	paginate.Total = int32(len(output))
	paginate.TotalRecords = int32(count)

	return output, paginate, nil
}

func (repo *Repository) GetUser(id int32) (output model.Users, err error) {
	result := repo.db.First(&model.Users{}, id).Scan(&output)

	if result.RowsAffected == 0 {
		return model.Users{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	if result.Error != nil {
		return model.Users{}, result.Error
	}

	return output, nil
}

func (repo *Repository) UpdateUser(input model.Users, id int32) (output model.Users, err error) {

	result := repo.db.Model(&input).Where("id", id).Updates(&input).Scan(&output)

	if result.RowsAffected == 0 {
		return model.Users{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	if result.Error != nil {
		return model.Users{}, result.Error
	}

	return output, nil
}

func (repo *Repository) DeleteUser(id int32) error {
	result := repo.db.Delete(&model.Users{}, id)

	if result.RowsAffected == 0 {
		return errors.New(gorm.ErrRecordNotFound.Error())
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
