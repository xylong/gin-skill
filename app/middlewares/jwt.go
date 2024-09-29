package middlewares

import (
	"gin-skill/app/common/response"
	"gin-skill/app/services"
	"gin-skill/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(context)
			context.Abort()
			return
		}
		tokenStr = tokenStr[len(services.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil {
			response.TokenFail(context)
			context.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(context)
			context.Abort()
			return
		}

		context.Set("token", token)
		context.Set("id", claims.Id)
	}
}
