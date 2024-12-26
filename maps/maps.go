package main

func Search(dict map[string]string, searchString string) string {
	val := dict[searchString]
	return val
}
