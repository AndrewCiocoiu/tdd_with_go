package iteration

import "strings"

// Will repeat the given string the given number of times
func Repeat(s string, times int) string {
	var repeated strings.Builder

	for i := 0; i < times; i++ {
		repeated.WriteString(s)
	}

	return repeated.String()
}
