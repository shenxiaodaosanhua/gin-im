package main

import (
	"my-im/src/controller/auth"
	"my-im/src/controller/index"
	"my-im/src/controller/message"
	"my-im/src/kernel/orm"
	"my-im/src/kernel/server"
	"my-im/src/middleware"
)

func main() {
	orm.InitDb()

	server.Ignite().Mount(
		"v1",
		auth.NewLoginController(),
	).Attach(
		middleware.NewAuthenticate(),
	).Mount(
		"v1",
		index.NewIndexController(),
		message.NewMessageController(),
	).Launch()
}
