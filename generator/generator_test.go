package generator

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://github.com/yusianglin11010/url-shortner"
	shortLink_1 := GenerateShortLink(initialLink_1)

	initialLink_2 := "https://google.com"
	shortLink_2 := GenerateShortLink(initialLink_2)

	assert.Equal(t, shortLink_1, "bVgqnjXw")
	assert.Equal(t, shortLink_2, "Lhr4BWAi")
}
