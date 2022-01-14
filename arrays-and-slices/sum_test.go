package arraysandslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertMessage := func(t testing.TB, expect, got int) {
		t.Helper()
		if expect != got {
			t.Errorf("\nexpected: %d\ngot: %d", expect, got)
		}
	}
	t.Run("sum up all the items in the passed collection", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers[:])
		expect := 15

		assertMessage(t, expect, got)
	})
}

func TestSumAllTails(t *testing.T) {
	assertMessage := func(t testing.TB, expect, got []int) {
		t.Helper()
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expected %v got %v", expect, got)
		}
	}

	t.Run("should return a slice containing the sums of all the items in the two slices except the first", func(t *testing.T) {
		expect := []int{29, 3}
		got := SumAllTails([]int{1, 8, 12, 9}, []int{6, 2, 1})
		assertMessage(t, expect, got)
	})

	t.Run("when one of the passed slices is empty, it's sum should be zero", func(t *testing.T) {
		expect := []int{0, 3}
		got := SumAllTails([]int{}, []int{6, 2, 1})

		assertMessage(t, expect, got)
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers[:])
	}
}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	//Output: 15
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 8, 12, 9}, []int{6, 2, 1})
	}
}

func ExampleSumAllTails() {
	collections, numbers := []int{1, 8, 12, 9}, []int{6, 2, 1}
	fmt.Println(SumAllTails(collections, numbers))
	// Output: [29 3]
}
