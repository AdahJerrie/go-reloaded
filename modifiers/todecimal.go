package modifiers

import (
	"regexp"
	"strconv"
)

func ReplaceTaggedNumbers(input string) string {

	// Step 1: Create regex patterns
	hexRe := regexp.MustCompile(`(?i)([0-9A-Fa-f]+)\s*\(hex\)`)
	binRe := regexp.MustCompile(`(?i)([01]+)\s*\(bin\)`)

	// Step 2: Replace HEX values
	input = hexRe.ReplaceAllStringFunc(input, func(match string) string {

		// Extract the number part
		parts := hexRe.FindStringSubmatch(match)

		// parts[1] is the number before (hex)
		numberStr := parts[1]

		// Convert hex → decimal
		decimalValue, err := strconv.ParseInt(numberStr, 16, 64)
		if err != nil {
			return match // if error, keep original
		}

		// Return decimal as string
		return strconv.FormatInt(decimalValue, 10)
	})

	// Step 3: Replace BINARY values
	input = binRe.ReplaceAllStringFunc(input, func(match string) string {

		parts := binRe.FindStringSubmatch(match)
		numberStr := parts[1]

		decimalValue, err := strconv.ParseInt(numberStr, 2, 64)
		if err != nil {
			return match
		}

		return strconv.FormatInt(decimalValue, 10)
	})

	return input
}
