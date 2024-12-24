package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const germanHelloPrefix = "Hallo, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"
const german = "German"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case german:
		prefix = germanHelloPrefix
	}
	return prefix + name
}

func World() string {
	return "Awesome!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
