package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == "French" {
		return "Bonjour, " + name
	}
	return englishHelloPrefix + name
}

func World() string {
	return "Awesome!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
