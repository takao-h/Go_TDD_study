package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
