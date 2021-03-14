package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cors struct {
}

func NewCors() *Cors {
	return &Cors{}
}

func (c *Cors) OnRequest(ctx *gin.Context) error {
	method := ctx.Request.Method

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}

	return nil
}

func (c *Cors) OnResponse(ctx *gin.Context) error {

	return nil
}
