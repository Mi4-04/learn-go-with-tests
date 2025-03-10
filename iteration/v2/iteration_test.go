package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 3)
	want := "aaa"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 4)
	fmt.Println(result)
	// Output: bbbb
}
