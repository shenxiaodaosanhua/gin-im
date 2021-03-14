package message

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"my-im/src/kernel/server"
	"my-im/src/kernel/ws"
	"my-im/src/model"
	"net/http"
)

type MessageController struct {
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

func (c *MessageController) Build(server *server.Server) {
	server.Handle("GET", "/ws", c.Ws)
}

func (c *MessageController) Ws(ctx *gin.Context) {
	client, err := ws.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	u, ok := ctx.Get("user")
	user := u.(model.UserClaim)
	if !ok {
		ctx.JSON(401, gin.H{"code": 401, "message": "未授权"})
	}

	ws.ClientMap.Store(&user, client)
}
