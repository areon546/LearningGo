package stringManip

import (
	"fmt"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	expected := "world"
	result := Cut("Hello, world")

	if expected != result {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestFields(t *testing.T) {
	input := "Test sentence"
	result := strings.Fields(input)[:2]
	expected := []string{"Test", "Sentence"}

	if expected != result {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func ExampleCut() {
	output := Cut("Hello, Jaqueline")
	fmt.Println(output)
	// Output: Jaqueline
}

func BenchmarkCut(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Cut("Hello, ")
	}
}
