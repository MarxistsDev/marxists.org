package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"marxists.org/models"
)

type AuthorController struct {
	Database *gorm.DB
}

// IndexHandler handles the root route.
// IndexHandler method
func (controller *AuthorController) AuthorById(c *gin.Context) {

	author_id, conversion_err := strconv.Atoi(c.Params.ByName("id"))

	if conversion_err != nil {
		c.HTML(http.StatusNotFound, "404.gohtml", gin.H{"Error": "Author not Found"})
		return
	}

	var author models.Author

	if err := controller.Database.Model(&models.Author{}).Joins("Glossary").First(&author, author_id).Error; err != nil {
		c.HTML(http.StatusNotFound, "404.gohtml", gin.H{"Error": "Author not Found"})
		return
	}

	//fmt.Println("Works:", author.Works)

	controller.Database.Model(&author).Preload("Works").First(&author, author_id)
	// Render the template with author data
	c.HTML(http.StatusOK, "author.gohtml", gin.H{"author": author})
}
