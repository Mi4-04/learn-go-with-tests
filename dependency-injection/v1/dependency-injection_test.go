package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "man")

	got := buffer.String()
	want := "Hello man"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
