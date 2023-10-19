package cloud

import (
	"errors"

	"github.com/google/uuid"
)

// Cloud represents a tag-cloud.
type Cloud struct {
	Id uuid.UUID `json:"id"`
	words []string
}

// New creates a new tag-cloud.
func New() *Cloud {
	return &Cloud{
		Id: uuid.New(),
		words: make([]string, 0),
	}
}

// Add adds a word to the tag-cloud.
func (s *Cloud) AddWord(word string) error {
	return errors.New("not implemented")
}
