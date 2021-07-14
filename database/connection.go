package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Con() *gorm.DB {
	dsn := "pilput:pilput31@tcp(127.0.0.1:3306)/learnfiber?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal cok")
	}
	return db
}
