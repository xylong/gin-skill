package main

import (
	user2 "gin-skill/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()

	router.GET("/users", func(context *gin.Context) {
		user := user2.NewUser(user2.WithName("tom")).
			Mutate(user2.WithGender(1))

		context.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	_ = router.Run("0.0.0.0:9000")
}
