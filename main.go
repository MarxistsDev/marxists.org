// main.go
package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"marxists.org/controllers"
	"marxists.org/models"
	"marxists.org/models/repository"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=marxists port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	router.SetFuncMap(template.FuncMap{"highlightSearch": repository.HighlightMatchingWords})
	router.StaticFile("style.css", "./www/styles/style.css")
	router.LoadHTMLGlob("views/*.gohtml")

	authorController := controllers.AuthorController{Repo: repository.AuthorRepository{Db: db}}
	searchController := controllers.SearchController{Repo: repository.SearchRepository{Db: db}}
	workController := controllers.WorkController{Repo: repository.WorkRepository{Db: db}}

	router.GET("/", controllers.IndexHandler)

	router.GET("/author/:id", authorController.AuthorById)

	router.GET("/search/:query", searchController.Search)

	router.GET("/work/:id/*ch", workController.Work)

	// Start the server
	router.Run("localhost:8080")
}
