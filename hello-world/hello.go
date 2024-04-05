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

	return fmt.Sprintf("%s, %s", createPrefix(language), name)
}

func createPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishPrefix
	case "French":
		return frenchPrefix
	default:
		return englishPrefix
	}
}
