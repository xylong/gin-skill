package main

import (
	"gin-skill/middlewares"
	user2 "gin-skill/models/user"
	"gin-skill/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(middlewares.RequestHandle())

	router.POST("/users", func(context *gin.Context) {
		//user := user2.NewUser(user2.WithName("tom")).
		//	Mutate(user2.WithGender(1))

		user := user2.NewUser()
		response.Result(context.ShouldBind(user)).Unwrap()

		context.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	_ = router.Run("0.0.0.0:9000")
}
