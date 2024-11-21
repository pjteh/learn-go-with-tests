package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestWorld(t *testing.T) {
	got := World()
	want := "Awesome!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
