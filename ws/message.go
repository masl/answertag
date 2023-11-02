package ws

// Message represents a message sent over the websocket for a new tag.
type Message struct {
	Tag     string `json:"tag"`
	CloudID string `json:"cloudId"`
}
