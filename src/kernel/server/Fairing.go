package server

import "github.com/gin-gonic/gin"

type Fairing interface {
	OnRequest(ctx *gin.Context) error
	OnResponse(ctx *gin.Context) error
}
