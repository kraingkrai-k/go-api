package config

import (
	"github.com/kraingkrai-k/go-api/src/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSQL() *gorm.DB {
	// TODO ?? migrate
	dsn := "root:1234@tcp(127.0.0.1:3307)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Table("users").AutoMigrate(&model.Users{})

	return db
}
