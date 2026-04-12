package main

import (
	"fmt"
	"strings"
)

func TextCase(text string) string {
	slice := strings.Fields(text)
	result := make([]string, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		word := slice[i]
		if i+1 < len(slice) {
			next := slice[i+1]

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
		}
		result = append(result, word)
		i++
		continue
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(TextCase("let it rain (up) again today"))
}
