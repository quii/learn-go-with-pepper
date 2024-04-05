package main

import "fmt"

func main() {
	fmt.Println(Hello(""))
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name
}
