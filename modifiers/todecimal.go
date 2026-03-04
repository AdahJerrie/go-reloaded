package modifiers

import (
	"regexp"
	"strconv"
)

func ToDecimal(input string) string {
	//give the regexp patterns for hex and bin
	hexexp := regexp.MustCompile(`(?i)(\b[a-fA-F0-9]+)\s*\(hex\)`)
	Binexp := regexp.MustCompile(`(?i)(\b[01]+)\s*\(hex\)`)

	//extract all the hexexp matches from input and put in input
	input = hexexp.ReplaceAllStringFunc(input, func(match string) string {

		//extract the desired part from the match
		part := hexexp.FindStringSubmatch(match)

		//Take only the part in the captured group
		numbr := part[1]

		//parse the number for int conversion to decimal
		hexconv, err := strconv.ParseInt(numbr, 16, 64)
		if err != nil {
			panic(err)
		}
		// return the decimal as string
		return strconv.FormatInt(hexconv, 10)
	})
	//For the binary to decimal
	//extract the bin from the string
	input = Binexp.ReplaceAllStringFunc(input, func(match string) string {

		//extract and narrow down the match to the captured group
		parts := Binexp.FindStringSubmatch(match)

		//get the desired part containing the hexa number
		numbr := parts[1]

		//convert to decimal
		Binconv, err := strconv.ParseInt(numbr, 16, 64)
		if err != nil {
			return match
		}

		return strconv.FormatInt(Binconv, 10)

	})
	return input

}
