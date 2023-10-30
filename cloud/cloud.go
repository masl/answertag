package cloud

import (
	"errors"

	"github.com/google/uuid"
)

// Cloud represents a tag-cloud.
type Cloud struct {
	ID uuid.UUID `json:"id"`
	words []string
}

// New creates a new tag-cloud.
func New() *Cloud {
	return &Cloud{
		ID: uuid.New(),
		words: make([]string, 0),
	}
}

// AddWord adds a word to the tag-cloud.
func (s *Cloud) AddWord(word string) error {
	return errors.New("not implemented")
}
