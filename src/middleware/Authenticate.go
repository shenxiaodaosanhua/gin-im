package middleware

import (
	"github.com/gin-gonic/gin"
	"my-im/src/service/auth"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")
		if authToken == "" {
			authToken = ctx.Query("Authorization")
		}

		user, err := auth.CheckAuth(authToken)
		if err != nil {
			ctx.JSON(402, gin.H{
				"code":    402,
				"message": err.Error(),
			})
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
