package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("hello REM")
	want := "Hello, REM"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}