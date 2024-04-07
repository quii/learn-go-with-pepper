package mocking

import (
	"bytes"
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	. "github.com/quii/pepper/matchers/io"
	"io"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	Expect[io.Reader](t, buffer).To(ToHaveOutputted(`3
2
1
Go!`))
	Expect(t, spySleeper.Calls).To(Equal(3))
}

func ToHaveOutputted(expected string) Matcher[io.Reader] {
	return HaveString(EqualTo(expected))
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
