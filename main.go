package main

import (
	"fmt"
	"strings"
)

const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(strings.ToLower(language)) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Bob", "spanish"))
}
