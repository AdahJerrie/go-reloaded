package modifiers

import (
	"fmt"
	"strconv"
	"strings"
)

func HexBinToDecimal(text string) string {
	field := strings.Fields(text)
	var result []string

	for i := 0; i < len(field); i++ {
		if i+1 < len(field) && field[i+1] == "(hex)" {
			converthex, err := strconv.ParseInt(field[i], 16, 64)
			if err != nil {
				fmt.Println("Error parsing hex")
				continue
			}
			field[i] = strconv.FormatInt(converthex, 10)
			result = append(result, field[i])
			i++
			continue
		}
		if i+1 < len(field) && field[i+1] == "(bin)" {
			convertbin, err := strconv.ParseInt(field[i], 2, 64)
			if err != nil {
				fmt.Println("Error parsing bin")
				continue
			}
			field[i] = strconv.FormatInt(convertbin, 10)
			result = append(result, field[i])
			i++
			continue
		}
		result = append(result, field[i])
	}
	return strings.Join(result, " ")
}
