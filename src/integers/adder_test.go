package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(1, 3)
	want := 4

	if got != want {
		t.Errorf("GOT: %d | WANT: %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 3)
	fmt.Println(sum)
}