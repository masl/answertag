package cloud

import (
	"strings"

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

// AddTag adds a new tag with the given name to the tag-cloud
// and keeps track of the number of occurrences.
func (c *Cloud) AddTag(name string) error {
	// lowercase the tag name
	name = strings.ToLower(name)

	// check if the tag already exists and increment the count
	for _, t := range c.Tags {
		if t.Name == name {
			t.Count++
			return nil
		}
	}

	// add the tag to the tag-cloud if it doesn't exist
	c.Tags = append(c.Tags, &Tag{
		Name:  name,
		Count: 1,
	})

	return nil
}

// AllTags returns all tags of the tag-cloud.
func (c *Cloud) AllTags() ([]*Tag, error) {
	return c.Tags, nil
}
