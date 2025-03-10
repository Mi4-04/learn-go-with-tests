package v2

import "testing"

func TestHello(t *testing.T) {
	got := Hello("man")
	want := "Hello, man"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
