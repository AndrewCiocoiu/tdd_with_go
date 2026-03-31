package iteration

// Repeat will repeat the given string 5 times
func Repeat(s string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated = repeated + s
	}

	return repeated
}
