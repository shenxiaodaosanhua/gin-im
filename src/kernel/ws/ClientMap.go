package ws

import (
	"github.com/gorilla/websocket"
	"my-im/src/model"
	"sync"
	"time"
)

var ClientMap *ClientMapStruct

type ClientMapStruct struct {
	data sync.Map
}

func init() {
	ClientMap = &ClientMapStruct{}
}

//保存上线用户
func (c *ClientMapStruct) Store(user *model.UserClaim, conn *websocket.Conn) {
	key := user.Mobile
	client := NewClient(conn, user)
	c.data.Store(key, client)
	go client.Ping(time.Second * 2)
	go client.ReadLoop()
	go client.HandlerLoop()
}

//删除下线用户
func (c *ClientMapStruct) Remove(user *model.UserClaim) {
	c.data.Delete(user.Mobile)
}

//向所有人发送消息
func (c *ClientMapStruct) SendAll(msg string) {
	c.data.Range(func(key, value interface{}) bool {
		value.(*Client).conn.WriteMessage(websocket.TextMessage, []byte(msg))
		return true
	})
}

//获取指定用户的客户端
func (c *ClientMapStruct) GetClient(mobile string) *Client {
	value, ok := c.data.Load(mobile)
	if !ok {
		return nil
	}

	return value.(*Client)
}

func (c *ClientMapStruct) GetOnline() *sync.Map {
	return &c.data
}
