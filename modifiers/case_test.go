package modifiers

import (
	"fmt"
	"testing"
)

func TestCase(t *testing.T) {
	input := "let us go (up) and play in the RAIN (low) by myself (cap)"
	expected := "let us GO and play in the rain by Myself"

	output := TextCase(input)
	fmt.Println(output)

	if output != expected {
		t.Log(output)
	}
}
