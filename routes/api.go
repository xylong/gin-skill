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
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来
	v1.Use(middlewares.LimitIP("200-H"))

	{
		// 注册
		v1.POST("/register", middlewares.LimitPerRoute("5-M"), v12.Auth.Register)
		// 登录
		v1.POST("/login", middlewares.LimitPerRoute("2-M"), v12.Auth.Login)

		authRouter := v1.Group("").Use(
			middlewares.JWTAuth(services.AppGuardName),
			middlewares.LimitIP("1000-H"),
		)
		{
			// 登出
			authRouter.POST("/logout", v12.Auth.Logout)

			// 个人信息
			authRouter.GET("/me", v12.User.Me)
		}
	}
}
