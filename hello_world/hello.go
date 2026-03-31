package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const helloSpanishPreifx = "Hola, "
const helloFrenchPrefix = "Bonjour, "
const helloRomanianPreifx = "Buna, "

func Hello(name string, language string) string {

	if name == "" {
		name = "Stranger"
	}

	prefix := helloEnglishPrefix

	switch language {
	case "spanish":
		prefix = helloSpanishPreifx
	case "french":
		prefix = helloFrenchPrefix
	case "romanian":
		prefix = helloRomanianPreifx
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
