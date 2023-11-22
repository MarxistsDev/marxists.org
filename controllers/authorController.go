package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
)

type AuthorController struct {
	Repo repository.AuthorRepository
}

func html404(c *gin.Context, errMsg string) {
	c.HTML(http.StatusNotFound, "404.gohtml", gin.H{"Error": errMsg})
}

// IndexHandler handles the root route.
// IndexHandler method
func (ctrl *AuthorController) AuthorById(c *gin.Context) {

	author_id, conversion_err := strconv.Atoi(c.Params.ByName("id"))

	if conversion_err != nil {
		html404(c, "The Parameter must be a number")
		return
	}

	author, err := ctrl.Repo.Get(author_id)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	c.HTML(http.StatusOK, "author.gohtml", gin.H{"author": author, "hasSearch": true})
}
