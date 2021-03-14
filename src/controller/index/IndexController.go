package index

import (
	"github.com/gin-gonic/gin"
	"my-im/src/kernel/server"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) Build(s *server.Server) {
	s.Handle("GET", "/", c.Index)
}

func (c *IndexController) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"data": "ok"})
}
