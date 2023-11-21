package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
)

type WorkController struct {
	Repo repository.WorkRepository
}

func (ctrl *WorkController) Work(c *gin.Context) {

	work_id, conversion_err := strconv.Atoi(c.Params.ByName("id"))

	if conversion_err != nil {
		html404(c, "The Parameter must be a number")
		return
	}

	work, err := ctrl.Repo.Get(work_id)
	if err != nil {
		html404(c, "Author not Found")
		return
	}

	c.HTML(http.StatusOK, "work.gohtml", gin.H{"work": work, "hasSearch": true})
}
