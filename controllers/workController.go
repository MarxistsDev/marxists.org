package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"marxists.org/models/repository"
	"marxists.org/views"
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

	if work_id < 0 {
		html404(c, "The Parameter must be a positive number")
		return
	}

	uwork_id := uint(work_id)

	work, err := ctrl.Repo.Get(uwork_id)
	if err != nil {
		fmt.Println("Work:", work.Title)
		html404(c, "Work not Found")
		return
	}

	// I assume that this means, that it got redirected to its Index
	if work.WorkID != uwork_id {
		for index, value := range work.Works {
			if value.WorkID == uwork_id {
				ch_id = index
			}
		}
	}

	fmt.Println(ch_err != nil, len(work.Works), ch_id)
	if ch_err != nil || len(work.Works) <= ch_id || ch_id < 0 {
		ch_id = 0
	}

	views.Base(views.Work(work, ch_id)).Render(c, c.Writer)
}
