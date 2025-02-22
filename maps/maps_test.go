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
		err := dictionary.Add(word, def)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, def)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, "this is another definition")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		newDefinition := "this is a new definition"
		dictionary := Dictionary{word: definition}
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}
		dictionary.Update(word, definition)
		assertDefinition(t, dictionary, word, newDefinition)
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
	got, err := dictionary.Search(word)
	want := def
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertHelper(t, got, want)
}
