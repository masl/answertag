// https://github.com/gorilla/websocket/blob/main/examples/chat/client.go
package ws

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/masl/answertag/cloud"
	"github.com/masl/answertag/tmpl"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	// cloudID is the id of the cloud this client is connected to.
	cloudID string
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("unexpected websocket close", "error", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		// unmarshal message
		var incomingData Message
		err = json.Unmarshal(message, &incomingData)
		if err != nil {
			slog.Error("unmarshal", "error", err)
			break
		}

		// get cloud by id from storage and add tag
		cld, err := c.hub.storage.ReadByID(incomingData.CloudID)
		if err != nil {
			slog.Error("read cloud by id from storage", "error", err)
			break
		}

		err = cld.AddTag(incomingData.Tag)
		if err != nil {
			slog.Error("add tag to cloud", "error", err)
			break
		}

		// update cloud in storage
		err = c.hub.storage.Update(cld)
		if err != nil {
			slog.Error("update cloud in storage", "error", err)
			break
		}

		// get all tags and write message with tags.html templates
		allTags, err := cld.AllTags()
		if err != nil {
			slog.Error("get all tags", "error", err)
			break
		}

		var responseBuffer ResponseBuffer
		if err := c.hub.tm.RenderTemplate(&responseBuffer, "tags", cloud.SupplementTagsWithFontSizes(allTags), &tmpl.RenderTemplateOptions{
			Layout: "tags",
		}); err != nil {
			slog.Error("execute tags.html template", "error", err)
			break
		}

		// broadcast the message
		c.hub.broadcast <- BroadcastMessage{
			CloudID: incomingData.CloudID,
			Message: responseBuffer.Bytes(),
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, cloudID string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), cloudID: cloudID}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
