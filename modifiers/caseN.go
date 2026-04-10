package main 

import (
	"strings"
	"fmt"
	"strconv"
	"log"
)

func CommandN(text string) string {
	slice := strings.Fields(text)
	result := make([]string, 0, len(slice))

	for i := 0; i < len(slice); i++{
		mark := slice[i]
		if strings.HasPrefix(mark, "(up,") {
			var num string 
			if strings.HasSuffix(mark, ")") {
				num = strings.TrimPrefix(strings.TrimSuffix(mark, ")"), "(up,")
				//result = append(result, slice[i])
				continue
			}
			if i+1 < len(slice) {
				num = strings.TrimSuffix(slice[i+1], ")")
				//result = append(result, slice[:i], slice[i+2:])
				continue
			}
			count,err := strconv.Atoi(num)
			if err != nil {
				log.Fatal()
			}

			start := i - count +1
			if start <0 {
				start = 0
			}

			for j := start; j<=i; j++ {
				slice[j] = strings.ToUpper(slice[j])
				result = append(result, slice[j])
				continue
			}

		}
		result = append(result, slice[i])
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(CommandN("let it rain (up, 2)"))
}