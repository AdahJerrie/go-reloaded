package modifiers

import "testing"

func TestHex(t *testing.T) {
	input := "It has been 10 (bin) years."
	expected := "It has been 2 years."

	outcome := ReplaceNum(input)

	if outcome != expected {
		t.Errorf("Actual-Result is %s, Expected-Result is %s\n", outcome, expected)
	}
}
