package v1

import (
	"gin-skill/app/common/response"
	"gin-skill/app/models"
	"gin-skill/app/requests"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	Auth = new(AuthController)
)

type AuthController struct {
	BaseApiController
}

// Register 注册
func (c *AuthController) Register(ctx *gin.Context) {
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

// Login 登录
func (c *AuthController) Login(ctx *gin.Context) {
	var (
		err  error
		req  request.LoginReq
		user *models.User
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		response.ValidateFail(ctx, request.GetErrorMsg(req, err))
		return
	}

	if err, user = services.AuthService.Login(req); err != nil {
		response.BusinessFail(ctx, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(ctx, err.Error())
			return
		}
		response.Success(ctx, tokenData)
	}
}

// Logout 登出
func (c *AuthController) Logout(ctx *gin.Context) {
	err := services.JwtService.JoinBlackList(ctx.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(ctx, "登出失败")
		return
	}
	response.Success(ctx, nil)
}
