package iteration

import (
	"fmt"
	"testing"
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
	// Output: aaaaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}
