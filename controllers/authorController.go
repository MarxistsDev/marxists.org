package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
	"marxists.org/views"
)

type AuthorController struct {
	Repo repository.AuthorRepository
}

func (ctrl *AuthorController) AuthorById(c *gin.Context) {

	author_id, conversion_err := strconv.Atoi(c.Param("id"))

	if conversion_err != nil {
		html404(c, "The Parameter must be a number")
		return
	}

	author, err := ctrl.Repo.Get(author_id)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	views.Base(views.Author(author)).Render(c, c.Writer)
}
