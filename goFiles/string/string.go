package stringManip

import (
	"fmt"
	"strings"
)

func stringManip() {
	input := "Test sentence"
	fmt.Println(strings.Fields(input))
}

func Cut(s string) string {
	after, _ := strings.CutPrefix(s, "Hello, ")
	return after
}
