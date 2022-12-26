package generator

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

const userID = "4940dd37-1b18-40c2-8970-32e8b60d532e"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://github.com/yusianglin11010/url-shortner"
	shortLink_1 := GenerateShortLink(initialLink_1, userID)

	initialLink_2 := "https://google.com"
	shortLink_2 := GenerateShortLink(initialLink_2, userID)

	assert.Equal(t, shortLink_1, "YiijuGty")
	assert.Equal(t, shortLink_2, "BW8hj7e3")
}
