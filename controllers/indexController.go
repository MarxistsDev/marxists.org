package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"marxists.org/models"
)

// IndexHandler method
func IndexHandler(c *gin.Context) {
	// Sample data for demonstration
	author := models.Author{
		AuthorID: 1,
		Name:     "JohnDoe",
		OldWorks: "sdfd",
	}

	// Render the template with author data
	c.HTML(http.StatusOK, "index.gohtml", gin.H{"author": author, "hasSearch": false})
}
