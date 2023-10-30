package inmemory

import (
	"errors"

	"github.com/masl/answertag/cloud"
)

var ErrAlreadyExists = errors.New("cloud already exists")

type Store struct {
	clouds map[string]*cloud.Cloud
	
	// TODO: add clouds
	tags []string
}

func New() *Store {
	return &Store{
		clouds: make(map[string]*cloud.Cloud),
	}
}

func (s *Store) Add(c *cloud.Cloud) error {
	if _, ok := s.clouds[c.Id.String()]; ok {
		return ErrAlreadyExists
	}

	s.clouds[c.Id.String()] = c

	return nil
}

func (s *Store) GetById(id string) (*cloud.Cloud, error) {
	return nil, errors.New("not implemented")
}

func (s *Store) AddTag(tag string) error {
	s.tags = append(s.tags, tag)
	return nil
}

func (s *Store) GetAllTags() ([]string, error) {
	return s.tags, nil
}
