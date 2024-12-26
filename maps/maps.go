package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("can't find word")
var ErrWordExists = errors.New("word already exists")

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
