package iteration

import "strings"

// Repeat will repeat the given string 5 times
func Repeat(s string) string {
	var repeated strings.Builder

	for i := 0; i < 5; i++ {
		repeated.WriteString(s)
	}

	return repeated.String()
}
