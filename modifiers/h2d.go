package modifiers

import (
	"strconv"
	"strings"
)

func ReplaceNum(input string) string {
	//split the input text into words
	text := strings.Fields(input)

	//create a loop to check for (hex) in each word
	for i := 1; i < len(text); i++ {
		if text[i] == "(hex)" && i > 0 {
			//Parse and convert the preivous word to an integer
			Hexconv, err := strconv.ParseInt(text[i-1], 16, 64)
			if err != nil {
				continue
			}
			//Replace the hexadecimal with its decimal equivalent
			text[i-1] = strconv.FormatInt(Hexconv, 10)

			//Remove the hex from the text
			text = append(text[:i], text[i+1:]...)
			i-- // Decrement i to account for the removed element
		} else if text[i] == "(bin)" && i > 0 {
			//Parse and convert the preivous word to an integer
			Binconv, err := strconv.ParseInt(text[i-1], 2, 64)
			if err != nil {
				continue
			}
			//Replace the binary with its decimal equivalent
			text[i-1] = strconv.FormatInt(Binconv, 10)

			//Remove the hex from the text
			text = append(text[:i], text[i+1:]...)
			i--
		}

	}
	return strings.Join(text, " ")
}
