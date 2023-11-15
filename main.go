// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"marxists.org/controllers"
)

func main() {
	//Establish connection
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Set up Gin router
	router := gin.Default()

	// Load HTML templates
	router.StaticFile("style.css", "./www/styles/style.css")
	router.LoadHTMLGlob("views/*.gohtml")

	// Define routes
	router.GET("/", controllers.IndexHandler)

	authorRoutes := router.Group("/author")
	{
		authorRoutes.GET("/", controllers.AuthorIndex(db))
		//authorRoutes.GET("/details", controllers.AuthorDetails)
	}

	// Start the server
	router.Run(":8080")
}
