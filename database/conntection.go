package database

import (
	"auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/user_auth?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
	DB = connection

	connection.AutoMigrate(&models.User{})
}
