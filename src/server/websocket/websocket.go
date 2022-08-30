package websocket

import (
	"net/http"
	"opengin/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func GetClient(ctx *gin.Context) *Client {
	token := ""
	secWebSocketProtocol := ctx.GetHeader("Sec-WebSocket-Protocol")

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{secWebSocketProtocol},
	}

	if secWebSocketProtocol == "" {
		return nil
	}

	token = secWebSocketProtocol

	if token == "" {
		token = ctx.Query("token")
	}

	clientId := ""
	claims, err := utils.ParseToken(token)

	if err != nil {
		return nil
	}

	clientId = claims.ClientId
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		return nil
	}

	hub := GetHub()
	client := &Client{
		Hub:     hub,
		Id:      clientId,
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}

	hub.RegisterClient(client)

	return client
}
