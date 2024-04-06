package iteration

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"testing"
)

func TestRepeat(t *testing.T) {
	Expect(t, Repeat("a")).To(Equal("aaaaa"))
}

func Repeat(s string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return repeated
}
