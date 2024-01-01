package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
	"marxists.org/views"
)

type SearchController struct {
	Repo repository.SearchRepository
}

func (ctrl *SearchController) Search(c *gin.Context) {
	query := c.Param("query")
	page_id, page_err := strconv.Atoi(strings.Replace(c.Param("page"), "/", "", 1))

	glossaries, err := ctrl.Repo.SearchGlossary(query)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	fmt.Println("Page id: ", page_id)
	if page_err != nil || page_id < 0 {
		page_id = 0
	}
	var upage_id uint = uint(page_id)
	articles, err := ctrl.Repo.SearchWork(query, &upage_id)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	fmt.Println("Article len", len(*articles))

	views.Base(views.Search(glossaries, articles, &query, &upage_id)).Render(c, c.Writer)
}
