package database

import (
	"log"
	"os"

	"github.com/cecep31/learnfiber/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConMysql() *gorm.DB {
	// dbname := os.Getenv("DB_NAME")
	dsn := "pilput:pilput31@tcp(127.0.0.1:3306)/learnfiber?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal cok")
	}
	return db
}

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConPsql() {
	dsn := "host=localhost user=pilput password=pilput31 dbname=learnfiber port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&migrations.Post{})

	DB = Dbinstance{
		Db: db,
	}
}
