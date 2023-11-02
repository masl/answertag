package cloud

import (
	"github.com/google/uuid"
)

// Cloud represents a tag-cloud.
type Cloud struct {
	ID   uuid.UUID `json:"id"`
	Tags []*Tag    `json:"tags"`
}

// New creates a new tag-cloud.
func New() *Cloud {
	return &Cloud{
		ID:   uuid.New(),
		Tags: make([]*Tag, 0),
	}
}

// AddTag adds a new tag to the tag-cloud.
func (c *Cloud) AddTag(t *Tag) error {
	c.Tags = append(c.Tags, t)
	return nil
}

// AllTags returns all tags of the tag-cloud.
func (c *Cloud) AllTags() ([]*Tag, error) {
	return c.Tags, nil
}
