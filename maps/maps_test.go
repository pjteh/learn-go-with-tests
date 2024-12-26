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

		if err == nil {
			t.Fatal("expected an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add single word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		def := "this is a test"
		dictionary.Add(word, def)
		assertDefinition(t, dictionary, word, def)
	})
}

func assertHelper(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func assertError(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, def string) {
	got, err := dictionary.Search("test")
	want := "this is a test"
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertHelper(t, got, want)
}
