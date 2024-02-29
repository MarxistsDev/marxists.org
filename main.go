// main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"marxists.org/initializers"
)

var DB *gorm.DB

func Init() {
	var err error

	if err = initializers.LoadEnv("./.env"); err != nil {
		log.Fatalf("Failed environment initialization: %v", err)
	}

	if err = initializers.ConnectDB(); err != nil {
		log.Fatalf("Failed database initialization: %v", err)
	}
}

func StartServer() {
	router := gin.Default()
	Routes(router)

	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func main() {
	Init()
	StartServer()
}
