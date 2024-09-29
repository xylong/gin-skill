package controllers

import (
	"gin-skill/app/common/response"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
)

func Me(ctx *gin.Context) {
	err, user := services.UserService.Me(ctx.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(ctx, err.Error())
		return
	}

	response.Success(ctx, user)
}
