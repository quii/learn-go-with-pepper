package main

import "fmt"

func main() {
	fmt.Println(Hello("", "english"))
}

const (
	spanishPrefix = "Hola"
	englishPrefix = "Hello"
	frenchPrefix  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishPrefix

	if language == "Spanish" {
		prefix = spanishPrefix
	}

	if language == "French" {
		prefix = frenchPrefix
	}

	return fmt.Sprintf("%s, %s", prefix, name)
}
