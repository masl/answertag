package ws

import (
	"bytes"
	"net/http"
)

// Message represents a message sent over the websocket for a new tag.
type Message struct {
	Tag     string `json:"tag"`
	CloudID string `json:"cloudId"`
}

// BroadcastMessage represents a broadcast message sent over the websocket
// to all clients including the tag cloud as html and the cloud id.
type BroadcastMessage struct {
	CloudID string `json:"cloudId"`
	Message []byte `json:"message"`
}

// ResponseBuffer is a buffer that implements http.ResponseWriter.
type ResponseBuffer struct {
	bytes.Buffer
}

func (rb *ResponseBuffer) Header() http.Header {
	return http.Header{}
}

func (rb *ResponseBuffer) WriteHeader(statusCode int) {
	// no-op
}
