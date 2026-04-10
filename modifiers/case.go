package modifiers

import "strings"

func TextCase(text string) string {
	slice := strings.Fields(text)
	var result []string

	for i := 0; i < len(slice); i++ {

		if i+1 < len(slice) && slice[i+1] == "(up)" {
			slice[i] = strings.ToUpper(slice[i])
			result = append(result, slice[i])
			i++
			continue
		}
		if i+1 < len(slice) && slice[i+1] == "(low)" {
			slice[i] = strings.ToLower(slice[i])
			result = append(result, slice[i])
			i++
			continue
		}
		if i+1 < len(slice) && slice[i+1] == "(cap)" {
			slice[i] = strings.ToUpper(slice[i][0:1]) + strings.ToLower(slice[i][1:])
			result = append(result, slice[i])
			i++
			continue
		}
		result = append(result, slice[i])
	}
	return strings.Join(result, " ")
}
