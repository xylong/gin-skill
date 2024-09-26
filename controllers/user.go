package controllers

import (
	"fmt"
	"gin-skill/dto"
	"gin-skill/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// GetUser 用户信息
func GetUser(ctx *gin.Context) (any, error) {
	var (
		req dto.UserInfoReq
	)

	if err := ctx.ShouldBindUri(&req); err != nil {
		return nil, fmt.Errorf("参数错误")
	}

	return service.GetSimpleUser(cast.ToInt64(req.Id))
}
