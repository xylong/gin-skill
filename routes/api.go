package routes

import (
	controllers2 "gin-skill/app/http/controllers"
	"gin-skill/app/http/middlewares"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册api路由
func RegisterAPIRoutes(router *gin.RouterGroup) {

	v1 := router.Group("/v1")
	{
		// 注册
		v1.POST("/register", controllers2.Register)
		// 登录
		v1.POST("/login", controllers2.Login)

		authRouter := v1.Group("").Use(middlewares.JWTAuth(services.AppGuardName))
		{
			// 登出
			authRouter.POST("/logout", controllers2.Logout)

			// 个人信息
			authRouter.GET("/me", controllers2.Me)
		}
	}
}
