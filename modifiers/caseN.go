package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(CommandN("let it rain (up,2) today"))
}

func CommandN(text string) string {
	field := strings.Fields(text)
	result := make([]string, 0, len(field))

	for i := 0; i < len(field); i++ {
		if strings.HasPrefix(field[i], "(up,") && strings.HasSuffix(field[i], ")") {
			numStr := strings.TrimPrefix(strings.TrimSuffix(field[i], ")"), "(up,")

			// if i+1 < len(field) && strings.HasSuffix(field[i+1], ")") {
			// 	numStr = strings.TrimSuffix(field[i+1], ")")
			// 	result = append(result, field[i])
			// 	i++
			// 	continue
			// }
			count, err := strconv.Atoi(numStr)
			if err != nil {
				result = append(result, field[i])
				continue
			}

			for j := len(result) - count; j < len(result); j++ {
				result[j] = strings.ToUpper(result[j])
				//result = append(result, field[j])
			}
			continue
		}
		result = append(result, field[i])
	}
	return strings.Join(result, " ")
}
