package bootstrap

import (
	"gin-skill/app/http/middlewares"
	"gin-skill/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// SetupRoute 路由初始化
func SetupRoute(engine *gin.Engine) *gin.Engine {
	// 注册全局中间件
	registerGlobalMiddleWare(engine)

	//  注册 API 路由
	routes.RegisterAPIRoutes(engine.Group("/api"))

	//  配置 404 路由
	setup404Handler(engine)

	return engine
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		middlewares.CustomRecovery(),
		middlewares.Cors(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
