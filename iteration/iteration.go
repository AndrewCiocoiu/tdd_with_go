package iteration

import "strings"

// Will repeat the given string the given number of times. Automatically removes spaces.
func Repeat(s string, times int) string {

	sanitizedString := strings.ReplaceAll(s, " ", "")

	var repeated strings.Builder

	for i := 0; i < times; i++ {
		repeated.WriteString(sanitizedString)
	}

	return repeated.String()
}
