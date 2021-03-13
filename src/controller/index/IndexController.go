package index

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": "ok"})
}
