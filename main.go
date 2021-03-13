package main

import (
	"github.com/gin-gonic/gin"
	"my-im/src/controller/auth"
	"my-im/src/controller/index"
	"my-im/src/controller/message"
	"my-im/src/kernel/ws"
	"my-im/src/middleware"
	"my-im/src/model"
)

func main() {
	model.InitDb()

	r := gin.New()

	loginController := auth.NewLoginController()
	r.POST("/login", loginController.Login)
	r.POST("/register", loginController.Register)

	r.Use(middleware.Auth())
	{
		indexController := index.NewIndexController()
		r.GET("/", indexController.Index)

		messageController := message.NewMessageController()
		r.GET("/ws", messageController.Ws)
		r.POST("/send", func(ctx *gin.Context) {
			msg, _ := ctx.GetPostForm("msg")
			ws.ClientMap.SendAll(msg)
		})
	}

	r.Run(":8080")
}
