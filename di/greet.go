package di

import (
	"fmt"
	"io"
)

func Greet(out io.Writer, name string) {
	fmt.Fprintf(out, "Hello, %s", name)
}
