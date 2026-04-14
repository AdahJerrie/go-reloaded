package modifiers

import (
	"regexp"
	"strings"
)

func Punc(text string) string {
	//handle space before punctuation
	regBefore := regexp.MustCompile(`\s+([.,;:!?]+)`)
	text = regBefore.ReplaceAllString(text, `$1`)

	//handle space after punctuation by creating two capture group
	regAfter := regexp.MustCompile(`([.,;:!?]+)([a-zA-Z])`)
	text = regAfter.ReplaceAllString(text, `$1 $2`)

	result := strings.TrimSpace(text)

	return result
}
