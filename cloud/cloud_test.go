package cloud_test

import (
	"testing"

	"github.com/masl/answertag/cloud"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("Returns a new cloud", func(t *testing.T) {
		c := cloud.New()

		assert.NotNil(t, c)
	})

	t.Run("Returns a new cloud with an unique ID", func(t *testing.T) {
		c1 := cloud.New()
		c2 := cloud.New()

		assert.NotNil(t, c1)
		assert.NotNil(t, c2)

		assert.NotEqual(t, c1.ID, c2.ID)
	})
}

func TestAddTag(t *testing.T) {
	t.Run("Adds a tag to the cloud", func(t *testing.T) {
		c := cloud.New()
		t1 := &cloud.Tag{Name: "lol", Count: 1}

		err := c.AddTag(t1)
		assert.NoError(t, err)

		assert.Equal(t, []*cloud.Tag{t1}, c.Tags)
	})

	t.Run("Adds multiple tags to the cloud", func(t *testing.T) {
		c := cloud.New()
		t1 := &cloud.Tag{Name: "lol", Count: 1}
		t2 := &cloud.Tag{Name: "rofl", Count: 2}

		err := c.AddTag(t1)
		assert.NoError(t, err)

		err = c.AddTag(t2)
		assert.NoError(t, err)

		assert.Equal(t, []*cloud.Tag{t1, t2}, c.Tags)
	})

	t.Run("Adds multiple tags with the same name to the cloud", func(t *testing.T) {
		c := cloud.New()
		t1 := &cloud.Tag{Name: "lol", Count: 1}

		tExpected := &cloud.Tag{Name: "lol", Count: 3}

		for i := 0; i < 3; i++ {
			err := c.AddTag(t1)
			assert.NoError(t, err)
		}

		assert.Equal(t, []*cloud.Tag{tExpected}, c.Tags)
	})
}

func TestAllTags(t *testing.T) {
	t.Run("Returns all tags of the cloud", func(t *testing.T) {
		c := cloud.New()
		t1 := &cloud.Tag{Name: "lol", Count: 1}
		t2 := &cloud.Tag{Name: "rofl", Count: 2}

		err := c.AddTag(t1)
		assert.NoError(t, err)

		err = c.AddTag(t2)
		assert.NoError(t, err)

		allTags, err := c.AllTags()
		assert.NoError(t, err)

		assert.Equal(t, []*cloud.Tag{t1, t2}, allTags)
	})

	t.Run("Returns an empty slice if the cloud has no tags", func(t *testing.T) {
		c := cloud.New()

		allTags, err := c.AllTags()
		assert.NoError(t, err)

		assert.Equal(t, []*cloud.Tag{}, allTags)
	})
}
