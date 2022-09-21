package config

import (
	"fmt"
	"github/exmaximum/passwordmanager/src/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DataBase *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connected to db")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.User{}, &models.Passwords{}, &models.PasswordNames{})
	DataBase = db
}
