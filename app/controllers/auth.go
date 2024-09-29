package controllers

import (
	"gin-skill/app/common/request"
	"gin-skill/app/common/response"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(ctx *gin.Context) {
	var (
		err error
		req request.RegisterReq
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(req.GetMessages(), err))
	}

	if err, user := services.AuthService.Register(req); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		response.Success(ctx, user)
	}
}
