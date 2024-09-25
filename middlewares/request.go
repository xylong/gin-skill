package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"msg": err,
				})
			}
		}()

		ctx.Next()
	}
}
