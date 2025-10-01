package database

import (
	"log"
	"taskmanager/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:123123@tcp(127.0.0.1:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối database: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Task{})

	DB = db
}
