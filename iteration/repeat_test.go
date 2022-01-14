package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertMessage := func(t testing.TB, expect, got string) {
		t.Helper()
		if expect != got {
			t.Errorf("\nexpected: %q\ngot: %q", expect, got)
		}
	}
	t.Run("loops through the array", func(t *testing.T) {
		expect := "aaaaa"
		got := Repeat("a", 5)

		assertMessage(t, expect, got)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("5", 12)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("r", 4))
	// Output: rrrr
}
