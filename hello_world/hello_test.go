package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people.", func(t *testing.T) {
		got := Hello("Andrei", "")
		want := "Hello, Andrei"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying stranger to people who don't provide their name.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Stranger"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet in Spanish", func(t *testing.T) {
		got := Hello("Andrew", "spanish")
		want := "Hola, Andrew"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet in French", func(t *testing.T) {
		got := Hello("Andrew", "french")
		want := "Bonjour, Andrew"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet in Romanian", func(t *testing.T) {
		got := Hello("Andrei", "romanian")
		want := "Buna, Andrei"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func ExampleHello() {
	output := Hello("Andrei", "romanian")
	fmt.Println(output)
	// Output: Buna, Andrei
}
