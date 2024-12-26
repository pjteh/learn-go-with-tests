package main

func Search(dict map[string]string, searchString string) string {
	val, _ := dict[searchString]
	return val
}
