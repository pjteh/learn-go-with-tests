package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("can't find word")

func (d Dictionary) Search(searchString string) (string, error) {
	val, ok := d[searchString]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}
