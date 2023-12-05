package database

import (
	"fmt"
	"log"
	"mangosteen/internal/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "192.168.1.15"
	user     = "root"
	password = "123456"
	dbname   = "mangosteen_dev"
	port     = 8888
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}

func Migrate() {
	DB.AutoMigrate(&model.User{}, &model.ValidationCode{})
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
}
