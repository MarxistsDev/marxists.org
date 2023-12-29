package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
)

type SearchController struct {
	Repo repository.SearchRepository
}

func (ctrl *SearchController) Search(c *gin.Context) {

	query := c.Param("query")

	glossaries, err := ctrl.Repo.SearchGlossary(query)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	articles, err := ctrl.Repo.SearchWork(query)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	c.HTML(http.StatusOK, "search.gohtml", gin.H{"glossaries": glossaries, "articles": articles, "query": query, "hasSearch": false})
}
