package arraysandslices

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {
	assertMessage := func(t testing.TB, expect, got int) {
		t.Helper()
		if expect != got {

			t.Errorf("expect %d got %d", expect, got)
		}
	}
	t.Run("multiplies the items contained in the collection", func(t *testing.T) {
		numbers := []int{2, 4, 8}
		expect := 64
		got := Product(numbers)
		assertMessage(t, expect, got)
	})

}

func BenchmarkProduct(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Product(numbers)
	}
}

func ExampleProduct() {
	numbers := []int{2, 4, 8}
	fmt.Println(Product(numbers))
	// Output: 64
}
