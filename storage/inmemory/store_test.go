package inmemory_test

import (
	"testing"

	"github.com/masl/answertag/cloud"
	"github.com/masl/answertag/storage"
	"github.com/masl/answertag/storage/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("Returns a new store", func(t *testing.T) {
		s := inmemory.New()

		assert.NotNil(t, s)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Returns nil if the game does not exist yet", func(t *testing.T) {
		s := inmemory.New()
		c := cloud.New()

		err := s.Create(c)
		assert.NoError(t, err)
	})

	t.Run("Returns an error if the cloud already exists", func(t *testing.T) {
		s := inmemory.New()
		c := cloud.New()

		err := s.Create(c)
		assert.NoError(t, err)

		err = s.Create(c)
		assert.Error(t, err)
		assert.Equal(t, storage.ErrAlreadyExists, err)
	})

	t.Run("Stores the cloud", func(t *testing.T) {
		s := inmemory.New()
		c := cloud.New()

		err := s.Create(c)
		assert.NoError(t, err)

		stored, err := s.ReadByID(c.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, c, stored)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Returns an error if the cloud does not exist", func(t *testing.T) {
		s := inmemory.New()
		c := cloud.New()

		err := s.Update(c)
		assert.Error(t, err)
		assert.Equal(t, storage.ErrNotFound, err)
	})

	t.Run("Updates the cloud", func(t *testing.T) {
		s := inmemory.New()
		c := cloud.New()

		err := s.Create(c)
		assert.NoError(t, err)

		c.AddTag(&cloud.Tag{
			Name:  "foo",
			Count: 1,
		})

		err = s.Update(c)
		assert.NoError(t, err)

		storedC, err := s.ReadByID(c.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, c, storedC)
	})
}

func TestReadByID(t *testing.T) {
	t.Run("Returns an error if the cloud does not exist", func(t *testing.T) {
		s := inmemory.New()

		_, err := s.ReadByID("goofy")
		assert.Error(t, err)
		assert.Equal(t, storage.ErrNotFound, err)
	})

	t.Run("Returns the cloud", func(t *testing.T) {
		s := inmemory.New()
		c := cloud.New()

		err := s.Create(c)
		assert.NoError(t, err)

		storedC, err := s.ReadByID(c.ID.String())
		assert.NoError(t, err)
		assert.Equal(t, c, storedC)
	})
}
