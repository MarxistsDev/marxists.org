package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"marxists.org/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Author{}, &models.Glossary{}, &models.Work{},
		&models.Collection{}, &models.Movement{}) //, &models.AuthorWork{})
}

