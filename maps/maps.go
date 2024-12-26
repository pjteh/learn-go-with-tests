package main

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("can't find word")
	ErrWordExists = DictionaryErr("word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(searchString string) (string, error) {
	val, ok := d[searchString]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(word, def string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}
	d[word] = def
	return nil
}

func (d Dictionary) Update(word, definition string) {
	_, err := d.Search(word)
	if err != nil {
		d.Add(word, definition)
	}
	d[word] = definition
}
