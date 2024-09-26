package routers

import (
	"gin-skill/controllers"
	"gin-skill/middlewares"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.New()

	// 注册
	router.POST("/register", middlewares.Wrapper(controllers.Register))

	return router
}
