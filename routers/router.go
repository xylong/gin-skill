package routers

import (
	"gin-skill/controllers"
	"gin-skill/middlewares"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares.RequestHandle())

	// 注册
	router.POST("/register", middlewares.Wrapper(controllers.Register))

	// 用户
	router.GET("/users/:id", middlewares.Wrapper(controllers.GetUser))

	return router
}
