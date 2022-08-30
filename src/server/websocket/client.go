package websocket

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = time.Second * 10
	pongWait   = 15 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	Id      string
	Socket  *websocket.Conn
	Message chan []byte
	Hub     *Hub
}

type MessageHandler func(c *Client, message []byte)

func (c *Client) Read(handler MessageHandler) {
	defer func() {
		c.Hub.UnregisterClient(c)
		c.Socket.Close()
	}()

	// c.Socket.SetReadDeadline(time.Now().Add(pongWait))

	for {
		messageType, message, readErr := c.Socket.ReadMessage()

		if readErr != nil || messageType == websocket.CloseMessage {
			fmt.Println(readErr.Error())
			break
		}

		handler(c, message)
	}
}

func (c *Client) Write() {
	defer func() {
		c.Hub.UnregisterClient(c)
		c.Socket.Close()
	}()

	// Ping ticker
	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case message, ok := <-c.Message:
			c.Socket.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.Socket.WriteMessage(websocket.TextMessage, message)

			if err != nil {
				fmt.Println(err.Error())
				return
			}
		case <-ticker.C:
			c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Socket.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
