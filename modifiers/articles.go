package modifiers

import (
	"strings"
)

func AtoAn(text string) string {
	field := strings.Fields(text)

	for i := 0; i < len(field); i++ {
		if i+1 < len(field) && strings.ContainsAny(field[i+1][:1], "aeiouAEIOUHh") {
			if field[i] == "a" {
				field[i] = "an"
			} else if field[i] == "A" {
				field[i] = "An"
			}
		}
	}
	return strings.Join(field, " ")
}
