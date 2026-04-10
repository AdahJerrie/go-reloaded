package modifiers

import (
	"fmt"
	"testing"
)

func TestHex(t *testing.T) {
	input := "It has been 10 (bin) years."
	expected := "It has been 2 years."

	outcome := HexBinToDecimal(input)
	fmt.Println(outcome)

	if outcome != expected {
		t.Errorf("Actual-Result is %s, Expected-Result is %s\n", outcome, expected)
	}
}
