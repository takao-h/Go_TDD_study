package iteration

import (
	"fmt"
	"testing"

	"golang.org/x/tools/go/expect"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 9)
	fmt.Printf(repeat)
}