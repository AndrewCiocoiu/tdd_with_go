package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people.", func(t *testing.T) {
		got := Hello("Andrei")
		want := "Hello, Andrei"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying stranger to people who don't provide their name.", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Stranger"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
