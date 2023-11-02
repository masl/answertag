package cloud

// Tag represents a single word-tag.
type Tag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
