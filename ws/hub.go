// https://github.com/gorilla/websocket/blob/main/examples/chat/hub.go
package ws

import (
	"text/template"

	"github.com/masl/answertag/storage"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan BroadcastMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Storage stores the tags.
	storage storage.Store

	// htmlTemplates includes the parsed templates.
	htmlTemplates *template.Template
}

func NewHub(store storage.Store, htmlTemplates *template.Template) *Hub {
	return &Hub{
		broadcast:     make(chan BroadcastMessage),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		storage:       store,
		htmlTemplates: htmlTemplates,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				// skip clients that are not connected to the same cloud
				if client.cloudID != message.CloudID {
					continue
				}
				select {
				case client.send <- message.Message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
