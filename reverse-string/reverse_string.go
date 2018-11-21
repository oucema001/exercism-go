package reverse

import (
	"strings"
)

func String(input string) string {
	if input == "" {
		return input
	}
	var b strings.Builder
	in := []rune(input)
	for i := len(in) - 1; i >= 0; i-- {
		b.WriteRune(in[i])
	}
	return b.String()
}
