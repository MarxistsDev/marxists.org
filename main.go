// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"marxists.org/controllers"
	"marxists.org/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Author{}, &models.Glossary{}, &models.Work{},
		&models.Article{}, &models.Collection{}, &models.Movement{}) //, &models.AuthorWork{})

	/*if err := db.SetupJoinTable(&models.Author{}, "Works", &models.AuthorWork{}); err != nil {
		println(models.AuthorWork{}.TableName(), err.Error())
		panic("Failed to setup join table")
	}*/

	router := gin.Default()
	router.StaticFile("style.css", "./www/styles/style.css")
	router.LoadHTMLGlob("views/*.gohtml")

	authorController := controllers.AuthorController{Database: db}

	router.GET("/", controllers.IndexHandler)

	authorRoutes := router.Group("/author")
	{
		authorRoutes.GET("/:id", authorController.AuthorById)
		//authorRoutes.GET("/details", controllers.AuthorDetails)
	}

	// Start the server
	router.Run("localhost:8080")
}