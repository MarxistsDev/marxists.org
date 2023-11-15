package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"marxists.org/models"
)

// IndexHandler handles the root route.
// IndexHandler method
func AuthorIndex(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var author models.Author
		if err := db.First(&author, 517).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Render the template with author data
		c.HTML(http.StatusOK, "author.gohtml", gin.H{"author": author})
	}
}
