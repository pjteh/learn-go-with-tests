package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'hello world' when no name is given", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
