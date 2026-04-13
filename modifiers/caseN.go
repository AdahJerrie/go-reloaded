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
		//detect (up,n)
		if strings.HasPrefix(field[i], "(up,") {
			numStr := strings.TrimPrefix(field[i], "(up,")

			//handle "(up," "n)"
			if i+1 < len(field) && !strings.HasSuffix(numStr, ")") {
				numStr = field[i+1]
				i++

			}
			numStr = strings.TrimSuffix(numStr, ")")

			count, err := strconv.Atoi(numStr)
			if err != nil {
				result = append(result, field[i])
				continue
			}

			start := len(result) - count
			if start < 0 {
				start = 0
			}

			for j := start; j < len(result); j++ {
				result[j] = strings.ToUpper(result[j])
			}
			continue
		}
		result = append(result, field[i])
	}
	return strings.Join(result, " ")
}
