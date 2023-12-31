// main.go
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"marxists.org/initializers"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	DB = initializers.DB
}

func main() {
	fmt.Print("env DSN :", os.Getenv("DSN"))
	router := gin.Default()

	Routes(router)
	// Start the server
	router.Run(os.Getenv("PORT"))
}
