package maps

import "testing"

func TestMaps(t *testing.T) {
	dictionary := Dictionary{"test": "something text"}

	got := dictionary.Search("test")
	want := "something text"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
