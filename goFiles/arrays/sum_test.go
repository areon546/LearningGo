package arraysAndSlices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("testing an array of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	expected := []int{6, 30}
	output := SumAll([]int{1, 2, 3}, []int{2, 4, 6, 8, 10})

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("got %d want %d", output, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSlices := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Sum Tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSlices(t, got, want)
	})

	t.Run("Sum tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSlices(t, got, want)
	})
}

// Shows a test case
func ExampleSumAll() {
	fmt.Println(SumAll([]int{1, 2, 3}))
	// Output: [6]
}
