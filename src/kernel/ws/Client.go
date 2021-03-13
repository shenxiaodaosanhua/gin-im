package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"my-im/src/model"
	"time"
)

type Client struct {
	conn      *websocket.Conn
	user      *model.UserClaim
	readChan  chan *Message
	closeChan chan struct{}
}

func NewClient(conn *websocket.Conn, user *model.UserClaim) *Client {
	return &Client{
		conn:      conn,
		user:      user,
		readChan:  make(chan *Message),
		closeChan: make(chan struct{}),
	}
}

func (c *Client) Ping(waitTime time.Duration) {
	for true {
		time.Sleep(waitTime)
		err := c.conn.WriteMessage(websocket.PingMessage, []byte("ping"))
		if err != nil {
			ClientMap.Remove(c.user)
			return
		}
	}
}

func (c *Client) ReadLoop() {
	for true {
		var msg *Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err.Error())
			c.conn.Close()
			ClientMap.Remove(c.user)
			c.closeChan <- struct{}{}
			break
		}

		c.readChan <- msg
	}
}

func (c *Client) HandlerLoop() {
loop:
	for true {
		select {
		case msg := <-c.readChan:
			c.HandlerMessage(msg)
		case <-c.closeChan:
			log.Println("已经关闭")
			break loop
		}
	}
}

func (c *Client) To(msg *Message) {
	toClient := ClientMap.GetClient(msg.To)
	if toClient == nil {
		fmt.Println("获取用户失败")
		return
	}
	toClient.Send(msg)
}

func (c *Client) HandlerMessage(msg *Message) {
	switch msg.Type {
	case MESSAGE_TEXT:
		c.To(msg)
	}
}

func (c *Client) Send(message *Message) {
	err := c.conn.WriteJSON(message)
	if err != nil {
		log.Fatal(err)
	}
}
