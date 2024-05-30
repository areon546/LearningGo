package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	output := Repeat("a", 5)

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}

}

func ExampleRepeat() {
	output := Repeat("a", 6)
	fmt.Println(output)
	// Output: aaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
