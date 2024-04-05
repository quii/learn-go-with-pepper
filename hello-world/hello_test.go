package main

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"testing"
)

func TestHello(t *testing.T) {
	Expect(t, Hello("")).To(EqualTo(`Hello, world`))
	Expect(t, Hello("Chris")).To(EqualTo(`Hello, Chris`))
}
