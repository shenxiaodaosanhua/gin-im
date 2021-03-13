package main

import (
	"github.com/gin-gonic/gin"
	"my-im/src/controller/auth"
	"my-im/src/model"
)

func main() {
	model.InitDb()

	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})

	loginController := auth.NewLoginController()
	r.POST("/login", loginController.Login)
	r.POST("/register", loginController.Register)

	r.Run(":8080")
}
