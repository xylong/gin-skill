package routes

import (
	v12 "gin-skill/app/http/controllers/api/v1"
	"gin-skill/app/http/middlewares"
	"gin-skill/app/services"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册api路由
func RegisterAPIRoutes(router *gin.RouterGroup) {

	v1 := router.Group("/v1")
	{
		// 注册
		v1.POST("/register", v12.Auth.Register)
		// 登录
		v1.POST("/login", v12.Auth.Login)

		authRouter := v1.Group("").Use(middlewares.JWTAuth(services.AppGuardName))
		{
			// 登出
			authRouter.POST("/logout", v12.Auth.Logout)

			// 个人信息
			authRouter.GET("/me", v12.User.Me)
		}
	}
}
