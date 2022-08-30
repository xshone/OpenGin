package websocket

import (
	"sync"
)

type Hub struct {
	Clients          map[string]*Client
	BroadcastMessage chan []byte
	Register         chan *Client
	Unregister       chan *Client
}

var hub *Hub
var once sync.Once

func GetHub() *Hub {
	once.Do(func() {
		hub = &Hub{
			Clients:          make(map[string]*Client),
			BroadcastMessage: make(chan []byte),
			Register:         make(chan *Client),
			Unregister:       make(chan *Client),
		}
	})

	return hub
}

func (h *Hub) RegisterClient(c *Client) {
	h.Register <- c
}

func (h *Hub) UnregisterClient(c *Client) {
	h.Unregister <- c
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.Id] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.Id]; ok {
				close(client.Message)
				delete(h.Clients, client.Id)
			}
		case message := <-h.BroadcastMessage:
			for clientId, client := range h.Clients {
				select {
				case client.Message <- message:
				default:
					close(client.Message)
					delete(h.Clients, clientId)
				}
			}
		}
	}
}
