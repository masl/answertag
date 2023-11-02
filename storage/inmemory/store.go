package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/masl/answertag/cloud"
)

var (
	ErrAlreadyExists = errors.New("cloud already exists")
	ErrNotFound      = errors.New("cloud not found")
)

// Store represents an in-memory storage.
type Store struct {
	// clouds maps cloud-IDs to clouds.
	clouds map[string]*cloud.Cloud
}

func New() *Store {
	return &Store{
		clouds: make(map[string]*cloud.Cloud),
	}
}

// Create creates a new tag-cloud.
func (s *Store) Create(c *cloud.Cloud) error {
	if _, ok := s.clouds[c.ID.String()]; ok {
		return ErrAlreadyExists
	}

	s.clouds[c.ID.String()] = c

	return nil
}

// Update updates an existing tag-cloud.
func (s *Store) Update(c *cloud.Cloud) error {
	if _, ok := s.clouds[c.ID.String()]; !ok {
		return ErrNotFound
	}

	s.clouds[c.ID.String()] = c

	return nil
}

// ReadById reads a tag-cloud by its ID.
func (s *Store) ReadById(id string) (*cloud.Cloud, error) {
	uuid.Parse(id)

	if cloud, ok := s.clouds[id]; ok {
		return cloud, nil
	}

	return nil, ErrNotFound
}
