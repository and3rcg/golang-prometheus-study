package internal

import (
	"fiber-prometheus-demo/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func StartDB() error {
	log.Println("Connecting to database...")
	dsn := "host=localhost user=postgres password=myPassword1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Database connection established!")
	Database = db

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations...")
	err = Database.AutoMigrate(&models.CardModel{})
	if err != nil {
		return err
	}

	return nil
}
