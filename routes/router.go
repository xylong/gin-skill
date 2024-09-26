package routes

import (
	controllers2 "gin-skill/app/controllers"
	middlewares2 "gin-skill/app/middlewares"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middlewares2.ErrorHandle())

	// 注册
	router.POST("/register", middlewares2.Wrapper(controllers2.Register))

	// 用户
	router.GET("/users/:id", middlewares2.Wrapper(controllers2.GetUser))

	return router
}
