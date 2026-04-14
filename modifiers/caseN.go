package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CommandN(input string) string {
	field := strings.Fields(input)

	for i := 0; i < len(field); i++ {
		if strings.HasPrefix(field[i], "(up,") {
			var num string
			if strings.HasSuffix(field[i], ")") {
				num = strings.TrimPrefix(strings.TrimSuffix(field[i], ")"), "(up,")
				field = append(field[:i], field[i+1:]...)
				i--
			}
			if i+1 < len(field) && strings.HasSuffix(field[i+1], ")") {
				num = strings.TrimSuffix(field[i+1], ")")
				field = append(field[:i], field[i+2:]...)
				i--
			}
			count, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error converting")
				continue
			}

			start := i - count + 1
			if start < 0 {
				start = 0
			}

			for j := start; j <= i; j++ {
				field[j] = strings.ToUpper(field[j])
			}
		}
	}
	return strings.Join(field, " ")
}

func main() {
	fmt.Println(CommandN("let it rain (up,2) today"))
}
