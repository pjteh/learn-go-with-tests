package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word search", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertHelper(t, got, want)
	})
	t.Run("unknown word search", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		_, err := dictionary.Search("what?")
		want := "can't find 'what?' in dictionary"

		if err == nil {
			t.Fatal("expected an error")
		}
		assertHelper(t, err.Error(), want)
	})
}

func assertHelper(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
