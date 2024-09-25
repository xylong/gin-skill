package main

import (
	"gin-skill/controllers"
	"gin-skill/middlewares"
	user2 "gin-skill/models/user"
	"gin-skill/pkg/response"
	"gin-skill/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(middlewares.RequestHandle())
	
	router.GET("/user", middlewares.Wrapper(controllers.GetUser))

	router.POST("/users", func(context *gin.Context) {
		//user := user2.NewUser(user2.WithName("tom")).
		//	Mutate(user2.WithGender(1))

		user := user2.NewUser()
		response.Result(context.ShouldBind(user)).Unwrap()

		response.OK(context)("", 0, response.Result(service.GetUserInfo(user.ID)).Unwrap())
	})

	_ = router.Run("0.0.0.0:9000")
}
