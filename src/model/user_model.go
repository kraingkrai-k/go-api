package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func (Users) TableName() string {
	return "users"
}

type Users struct {
	ID        int32     `json:"id" gorm:"primaryKey"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	Age       int32     `json:"age"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
}

func ValidateUsersSortKey(key string) error {
	if key == "" {
		return nil
	}

	fields := []string{"id", "age"}
	for _, field := range fields {
		if field == key {
			return nil
		}
	}

	return fmt.Errorf("sortKey is oneof %s", strings.Join(fields, ", "))
}

func (u *Users) Filter(filter Filter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if filter.Offset != 0 {
			db.Offset(int(filter.Offset))
		}
		if filter.Limit != 0 {
			db.Limit(int(filter.Limit))
		}

		sortKey := "id"
		sortBy := "asc"

		if filter.SortKey != "" {
			sortKey = filter.SortKey
		}

		if filter.SortBy != "" {
			sortBy = filter.SortBy
			sort := fmt.Sprintf("%s %s", sortKey, sortBy)
			db.Order(sort)
		}

		return db
	}
}

func (u *Users) AfterUpdate(tx *gorm.DB) (err error) {
	if u.ID == 2 {
		return errors.New("no !!, He is a GOD")
	}
	return nil
}

type UserOutput struct {
	ID        int32  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (x *Users) Validate() error {
	if err := validate.Struct(x); err != nil {
		return err
	}
	return nil
}
