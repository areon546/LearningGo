package stringManip

import (
	"strings"
)

func Cut(s string) string {
	after, _ := strings.CutPrefix(s, "Hello, ")
	return after
}
