package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
	"marxists.org/views"
)

type AuthorController struct {
	Repo repository.AuthorRepository
}

func (ctrl *AuthorController) AuthorById(c *gin.Context) {
	author_id, conversion_err := strconv.Atoi(c.Param("id"))
	page_id, page_err := strconv.Atoi(strings.Replace(c.Param("page"), "/", "", 1))

	if conversion_err != nil {
		html404(c, "The Parameter must be a number")
		return
	}

	fmt.Println("Page id: ", page_id)
	if page_err != nil || page_id < 0 {
		page_id = 0
	}
	var upage_id uint = uint(page_id)
	author, err := ctrl.Repo.Get(author_id, &page_id)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	views.Base(views.Author(author, &upage_id)).Render(c, c.Writer)
}
