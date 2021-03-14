package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	*gin.Engine
	groupRoute *gin.RouterGroup
}

func Ignite() *Server {
	return &Server{
		Engine: gin.New(),
	}
}

func (s *Server) Launch() {
	err := s.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Server {
	s.groupRoute.Handle(httpMethod, relativePath, handlers...)
	return s
}

//挂载控制器
func (s *Server) Mount(group string, handlers ...IHandler) *Server {
	s.groupRoute = s.Group(group)

	for _, handler := range handlers {
		handler.Build(s)
	}

	return s
}

//添加中间件
func (s *Server) Attach(f Fairing) *Server {
	s.Use(func(ctx *gin.Context) {
		err := f.OnRequest(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
		err = f.OnResponse(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
	})

	return s
}
