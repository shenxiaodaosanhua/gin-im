package ws

import (
	"bytes"
	"encoding/json"
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

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

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
	defer c.conn.Close()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for true {
		var msg *Message
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.conn.Close()
				ClientMap.Remove(c.user)
				c.closeChan <- struct{}{}
			}
			fmt.Println(err.Error())
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		err = json.Unmarshal(message, &msg)
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
