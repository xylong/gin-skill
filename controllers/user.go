package controllers

import (
	user2 "gin-skill/models/user"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) (any, error) {
	user := user2.NewUser(user2.WithName("summer"))
	return user, nil
}
