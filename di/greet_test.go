package di

import (
	"bytes"
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	. "github.com/quii/pepper/matchers/io"
	"io"
	"testing"
)

func TestGreeter(t *testing.T) {
	buffer := &bytes.Buffer{}
	Greet(buffer, "Chris")

	Expect[io.Reader](t, buffer).To(HaveString(EqualTo("Hello, Chris")))
}
