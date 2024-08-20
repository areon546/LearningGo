package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("2 + 2 = 4", func(t *testing.T) {
		result := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, expected, result)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, expected, result int) {
	t.Helper()
	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}
