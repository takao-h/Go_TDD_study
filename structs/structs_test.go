package structs

import "testing"

func TestPerimeter(t *testing.T) {
	width := 10.0
	height := 10.0

	got := Perimeter(width, height)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area()
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
