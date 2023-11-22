package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
)

type WorkController struct {
	Repo repository.WorkRepository
}

func (ctrl *WorkController) Work(c *gin.Context) {

	work_id, conversion_err := strconv.Atoi(c.Param("id"))
	ch_id, ch_err := strconv.Atoi(strings.Replace(c.Param("ch"), "/", "", 1))

	if conversion_err != nil {
		html404(c, "The Parameter must be a number")
		return
	}

	work, err := ctrl.Repo.Get(work_id)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	fmt.Println(ch_err != nil, len(work.Articles), ch_id)
	if ch_err != nil || len(work.Articles) <= ch_id || ch_id < 0 {
		ch_id = 0
	}

	c.HTML(http.StatusOK, "work.gohtml", gin.H{"work": work, "index": ch_id, "hasSearch": true})
}
