package v1

import (
	"gin-skill/app/common/response"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
)

var (
	User = new(UserController)
)

type UserController struct{}

// Me 个人信息
func (u *UserController) Me(ctx *gin.Context) {
	err, user := services.UserService.Me(ctx.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(ctx, err.Error())
		return
	}

	response.Success(ctx, user)
}
