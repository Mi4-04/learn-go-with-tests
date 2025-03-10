package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "something text"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "something text"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test1", "something text 1")

	want := "something text 1"
	got, err := dictionary.Search("test1")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
