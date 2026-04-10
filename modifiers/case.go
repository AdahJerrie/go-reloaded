package modifiers

import (
	"strings"
)

func TextCase(text string) string {
	slice := strings.Fields(text)
	result := make([]string, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		next := slice[i+1]
		word := slice[i]

		switch next {
		case "(up)":
			word = strings.ToUpper(word)
		case "(low)":
			word = strings.ToLower(word)
		case "(cap)":
			word = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		default:
			result = append(result, word)
			continue
		}
		result = append(result, word)
		i++
		continue
	}
	return strings.Join(result, " ")
}
