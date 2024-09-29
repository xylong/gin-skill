package routes

import (
	controllers2 "gin-skill/app/http/controllers"
	"gin-skill/app/http/middlewares"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	// 注册
	router.POST("/register", controllers2.Register)
	// 登录
	router.POST("/login", controllers2.Login)

	authRouter := router.Group("").Use(middlewares.JWTAuth(services.AppGuardName))
	{
		// 登出
		authRouter.POST("/logout", controllers2.Logout)

		// 个人信息
		authRouter.GET("/me", controllers2.Me)
	}
}
