package middleware

import (
	"github.com/gin-gonic/gin"
	"my-im/src/service/auth"
)

type Authenticate struct {
}

func NewAuthenticate() *Authenticate {
	return &Authenticate{}
}

func (a *Authenticate) OnRequest(ctx *gin.Context) error {
	authToken := ctx.GetHeader("Authorization")
	if authToken == "" {
		authToken = ctx.Query("Authorization")
	}

	user, err := auth.CheckAuth(authToken)
	if err != nil {
		return err
	}

	ctx.Set("user", user)
	return nil
}
