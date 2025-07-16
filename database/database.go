package database

import (
	"fmt"
	"log"
	"backend-go/config"
	"backend-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "")
	dbPort := config.GetEnv("DB_PORT", "")
	dbName := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("faild to connect to database :(", err)
	}

	fmt.Println("Database Connected successfully :)")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("aild to connect to database :(")
	}

	fmt.Println("Database Connected successfully :)")
}