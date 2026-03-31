package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const defaultGreeting = "Hello, Stranger"

func Hello(name string) string {

	if name == "" {
		return defaultGreeting
	}

	return helloEnglishPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
