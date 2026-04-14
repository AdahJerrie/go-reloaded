package main

import (
	"go-reloaded/modifiers"
	"log"
	"os"
	"strings"
)

func ProcessText(text string) string {
	input := modifiers.HexBinToDecimal(text)
	input = modifiers.AtoAn(input)
	input = modifiers.TextCase(input)
	input = modifiers.CommandN(input)
	input = modifiers.Punc(input)
	input = modifiers.SingleQuote(input)

	return input
}

func Newline(text string) string {
	word := strings.Split(text, "\n")

	for i, line := range word {
		word[i] = ProcessText(line)
	}
	return strings.Join(word, "\n")
}

func main() {
	inputfile := os.Args[1]
	outputfile := os.Args[2]

	file, err := os.ReadFile(inputfile)
	if err != nil {
		log.Fatal(err)
	}

	result := Newline(string(file))

	err = os.WriteFile(outputfile, []byte(result), 0644)
}
