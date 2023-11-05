package ws

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
