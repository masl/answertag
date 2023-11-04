package cloud_test

import (
	"testing"

	"github.com/masl/answertag/cloud"
	"github.com/stretchr/testify/assert"
)

func TestSupplementTagsWithFontSizes(t *testing.T) {
	t.Run("Returns a list of tags with font sizes", func(t *testing.T) {
		tags := []*cloud.Tag{
			{Name: "lol", Count: 1},
			{Name: "rofl", Count: 4},
			{Name: "lel", Count: 8},
		}

		tagsWithFontSizes := cloud.SupplementTagsWithFontSizes(tags)

		assert.Equal(t, []cloud.TagWithFontSize{
			{Tag: cloud.Tag{Name: "lol", Count: 1}, FontSize: 20},
			{Tag: cloud.Tag{Name: "rofl", Count: 4}, FontSize: 45},
			{Tag: cloud.Tag{Name: "lel", Count: 8}, FontSize: 80},
		}, tagsWithFontSizes)
	})

	t.Run("Returns nil if the given list of tags is empty", func(t *testing.T) {
		tagsWithFontSizes := cloud.SupplementTagsWithFontSizes([]*cloud.Tag{})

		assert.Nil(t, tagsWithFontSizes)
	})

	t.Run("Returns a list of tags with the same font size if all tags have the same count", func(t *testing.T) {
		tags := []*cloud.Tag{
			{Name: "lol", Count: 1},
			{Name: "rofl", Count: 1},
			{Name: "lel", Count: 1},
		}

		tagsWithFontSizes := cloud.SupplementTagsWithFontSizes(tags)

		for _, tagWithFontSize := range tagsWithFontSizes {
			assert.Equal(t, 40, tagWithFontSize.FontSize)
		}
	})
}
