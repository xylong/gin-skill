package controllers

import (
	"fmt"
	"gin-skill/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// GetUser 用户信息
func GetUser(ctx *gin.Context) (any, error) {
	//id := ctx.Param("id")
	id := &struct {
		Id int `uri:"id" binding:"required,gt=0"`
	}{}

	if err := ctx.ShouldBindUri(id); err != nil {
		return nil, fmt.Errorf("参数错误")
	}

	return service.GetUserById(cast.ToInt64(id.Id))
}
