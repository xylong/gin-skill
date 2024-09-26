package controllers

import "github.com/gin-gonic/gin"

// Register 注册
func Register(ctx *gin.Context) (any, error) {
	return gin.H{"a": "app"}, nil
}
