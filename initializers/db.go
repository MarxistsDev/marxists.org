package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"marxists.org/models"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return fmt.Errorf("DSN is not set")
	}

	loggerConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	config := &gorm.Config{
		PrepareStmt: true,
		Logger:      loggerConfig,
	}

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = DB.AutoMigrate(&models.Author{}, &models.Glossary{}, &models.Work{},
		&models.Collection{}, &models.Movement{}) //, &models.AuthorWork{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate database models: %w", err)
	}

	return nil
}
