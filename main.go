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

func init() {
	var err error

	if initializers.LoadEnv(".env"); err != nil {
		log.Fatalf("Failed environment initialization: %v", err)
	}

	if initializers.ConnectDB(); err != nil {
		log.Fatalf("Failed database initialization: %v", err)
	}
}

func main() {
	router := gin.Default()

	Routes(router)

	router.Run(os.Getenv("PORT"))
}
