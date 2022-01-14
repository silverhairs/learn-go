package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertMessage := func(t testing.TB, expect, got int) {
		t.Helper()
		if expect != got {
			t.Errorf("\nexpected: %d\n got: %d", expect, got)
		}
	}

	t.Run("returns the sum of the two passed numbers", func(t *testing.T) {
		expect := 20
		got := Adder(15, 5)
		assertMessage(t, expect, got)
	})
}

func ExampleAdder() {
	result := Adder(5, 10)
	fmt.Println(result)
	// Output: 15
}
