package modifiers

import (
	"regexp"
)

func SingleQuote(text string) string {
	reg := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = reg.ReplaceAllString(text, `'$1'`)

	return text
}
