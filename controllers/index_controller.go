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

	/*
			const authors = [
		  {
		    id: 1,
		    author: "Lenin",
		    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
		  },
		  {
		    id: 2,
		    author: "Marx",
		    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
		  },
		  {
		    id: 3,
		    author: "Engel",
		    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
		  }
		];
	*/

	// Render the template with author data
	c.HTML(http.StatusOK, "index.gohtml", gin.H{"author": author, "hasSearch": false})
}
