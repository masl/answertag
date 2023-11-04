package cloud

type TagWithFontSize struct {
	Tag
	FontSize int
}

const (
	minFontSize      = 20
	maxFontSize      = 80
	fallbackFontSize = 40
)

// SupplementTagsWithFontSizes supplements the given tags with font sizes based on their count.
func SupplementTagsWithFontSizes(tags []*Tag) []TagWithFontSize {
	if len(tags) == 0 {
		return nil
	}

	minCount, maxCount := findMinAndMaxCount(tags)
	countRange := maxCount - minCount

	// iterate over tags and calculate font size
	result := make([]TagWithFontSize, len(tags))
	for i, tag := range tags {
		// if all tags have the same count, use fallback font size
		if countRange == 0 {
			result[i] = TagWithFontSize{
				Tag:      *tag,
				FontSize: fallbackFontSize,
			}

			continue
		}

		// calculate relative count and font size
		relCount := float64(tag.Count-minCount) / float64(countRange)
		result[i] = TagWithFontSize{
			Tag:      *tag,
			FontSize: calcFontSize(relCount),
		}
	}

	return result
}

// calcFontSize calculates the font size based on the relative count.
func calcFontSize(relCount float64) int {
	return minFontSize + int(relCount*float64(maxFontSize-minFontSize))
}

// findMinAndMaxCount finds the minimum and maximum count in the given tags.
func findMinAndMaxCount(tags []*Tag) (int, int) {
	min := tags[0].Count
	max := tags[0].Count

	for _, tag := range tags {
		if tag.Count < min {
			min = tag.Count
		}

		if tag.Count > max {
			max = tag.Count
		}
	}

	return min, max
}
