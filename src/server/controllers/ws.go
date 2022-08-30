package controllers

import (
	"opengin/server/utils"
	"opengin/server/websocket"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Chat(ctx *gin.Context) {
	client := websocket.GetClient(ctx)

	if client == nil {
		return
	}

	go client.Read(func(c *websocket.Client, message []byte) {
		broadcastMessage := "[Broadcast]: " + string(message)
		c.Hub.BroadcastMessage <- []byte(broadcastMessage)
		privateMessage := "[" + c.Id + "]_" + "[Private]: " + string(message)
		c.Message <- []byte(privateMessage)
	})
	go client.Write()
}

func (c *Controller) GetMessage(ctx *gin.Context) {
	client := websocket.GetClient(ctx)

	if client == nil {
		return
	}

	mq := utils.NewRabbitMqHandler(&utils.RabbitMqOptions{
		ExchangeName: "gin",
		QueueName:    "test",
		RoutingKey:   "",
	})

	go client.Read(func(c *websocket.Client, message []byte) {
		if strings.ToLower(string(message)) == "stop" {
			mq.Close()
		}
	})
	go client.Write()

	mq.StartConsuming(func(msg []byte) {
		privateMessage := "[" + client.Id + "]_" + "[Private]: " + string(msg)
		client.Message <- []byte(privateMessage)
	})
}
