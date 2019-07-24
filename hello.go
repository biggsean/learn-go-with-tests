package main

import "fmt"

const (
	spanish            string = "Spanish"
	englishHelloPrefix string = "Hello, "
	spanishHelloPrefix string = "Hola, "
)

// Hello says hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
