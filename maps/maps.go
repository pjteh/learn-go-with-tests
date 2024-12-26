package main

import "fmt"

type Dictionary map[string]string

func (d Dictionary) Search(searchString string) (string, error) {
	val, ok := d[searchString]
	if !ok {
		return "", fmt.Errorf("can't find '%s' in dictionary", searchString)
	}
	return val, nil
}
