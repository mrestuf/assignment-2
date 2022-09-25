package controllers

import (
	"assignment2/httpserver/controllers/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonRespnse(ctx *gin.Context, resp *views.Response) {
	ctx.JSON(resp.Status, resp)
}
