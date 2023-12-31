package controllers

import (
	"github.com/gin-gonic/gin"
	"marxists.org/views"
)

func IndexHandler(c *gin.Context) {
	views.Base(views.Index()).Render(c, c.Writer)
}

func html404(c *gin.Context, errMsg string) {
	views.Error404(errMsg).Render(c, c.Writer)
}
