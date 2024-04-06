package integers

import (
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdder(t *testing.T) {
	Expect(t, Add(2, 2)).To(Equal(4))
}
