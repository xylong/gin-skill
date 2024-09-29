package routes

import (
	"gin-skill/app/controllers"
	"gin-skill/app/middlewares"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	// 注册
	router.POST("/register", controllers.Register)
	// 登录
	router.POST("/login", controllers.Login)

	authRouter := router.Group("").Use(middlewares.JWTAuth(services.AppGuardName))
	{
		// 个人信息
		authRouter.GET("/me", controllers.Me)
	}
}
