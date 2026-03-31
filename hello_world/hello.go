package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const helloSpanishPreifx = "Hola, "
const helloFrenchPrefix = "Bonjour, "
const defaultGreeting = "Hello, Stranger"

func Hello(name string, language string) string {

	if name == "" {
		return defaultGreeting
	}

	if language == "spanish" {
		return helloSpanishPreifx + name
	}

	if language == "french" {
		return helloFrenchPrefix + name
	}

	return helloEnglishPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
