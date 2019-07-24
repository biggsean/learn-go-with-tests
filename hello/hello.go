package main

import "fmt"

const (
	french             string = "French"
	spanish            string = "Spanish"
	englishHelloPrefix string = "Hello, "
	frenchHelloPrefix  string = "Bonjour, "
	spanishHelloPrefix string = "Hola, "
)

// Hello says hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
