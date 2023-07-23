package map

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is jest a test"}

	got := Search(dictionary, "test")
	want := "this is jest a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}