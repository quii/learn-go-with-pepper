package main

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"testing"
)

func TestHello(t *testing.T) {
	Expect(t, Hello("", "English")).To(EqualTo(`Hello, world`))
	Expect(t, Hello("Chris", "English")).To(EqualTo(`Hello, Chris`))
	Expect(t, Hello("Elodie", "Spanish")).To(EqualTo(`Hola, Elodie`))
	Expect(t, Hello("Milo", "French")).To(EqualTo(`Bonjour, Milo`))
}
