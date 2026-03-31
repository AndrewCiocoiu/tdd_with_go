package main

import "fmt"

const (
	spanish  = "spanish"
	romanian = "romanian"
	french   = "french"
	english  = "english"

	englishPrefix  = "Hello, "
	spanishPrefix  = "Hola, "
	frenchPrefix   = "Bonjour, "
	romanianPrefix = "Buna, "

	defaultName = "Stranger"
)

func Hello(name string, language string) string {

	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	case romanian:
		prefix = romanianPrefix
	case english:
		prefix = englishPrefix
	default:
		prefix = englishPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
