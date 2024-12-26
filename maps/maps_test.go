package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	got := dictionary.Search("test")
	want := "this is a test"
	assertHelper(t, got, want)
}

func assertHelper(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
