package main

type Dictionary map[string]string

func (d Dictionary) Search(searchString string) string {
	return d[searchString]
}
