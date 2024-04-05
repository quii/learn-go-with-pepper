package main

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"testing"
)

func TestHello(t *testing.T) {
	Expect(t, Hello("", "English")).To(Equal(`Hello, world`))
	Expect(t, Hello("Chris", "English")).To(Equal(`Hello, Chris`))
	Expect(t, Hello("Elodie", "Spanish")).To(Equal(`Hola, Elodie`))
	Expect(t, Hello("Milo", "French")).To(Equal(`Bonjour, Milo`))
}
