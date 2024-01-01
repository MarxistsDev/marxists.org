// main.go
package main

import (
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
	router := gin.Default()

	Routes(router)

	router.Run(os.Getenv("PORT"))
}
