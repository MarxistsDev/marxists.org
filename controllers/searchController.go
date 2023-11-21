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

	query := c.Params.ByName("query")

	articles, err := ctrl.Repo.SearchArticle(query)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	c.HTML(http.StatusOK, "search.gohtml", gin.H{"articles": articles, "query": query, "hasSearch": true})
}
