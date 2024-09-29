package routes

import (
	"gin-skill/app/controllers"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	// 注册
	router.POST("/register", controllers.Register)
	// 登录
	router.POST("/login", controllers.Login)
}
