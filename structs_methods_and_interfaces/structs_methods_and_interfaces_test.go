package structs_methods_and_interfaces

import (
	"fmt"
	. "github.com/quii/pepper"
	"math"
	"testing"
)

func TestShapes(t *testing.T) {
	Expect[Shape](t, Rectangle{10.0, 10.0}).To(
		HavePerimeter(40.0).And(HaveArea(100.0)),
	)
	Expect[Shape](t, EquilateralTriangle{Width: 10.0}).To(
		HaveArea(43.3012701892).And(HavePerimeter(30.0)),
	)
}

func HaveArea(want float64) Matcher[Shape] {
	return func(got Shape) MatchResult {
		return MatchResult{
			Description: fmt.Sprintf("to have area of %.5f", want),
			Matches:     almostEqual(got.Area(), want),
			But:         fmt.Sprintf("it had area of %.5f", got.Area()),
			SubjectName: "Shape",
		}
	}
}

func HavePerimeter(want float64) Matcher[Shape] {
	return func(got Shape) MatchResult {
		return MatchResult{
			Description: fmt.Sprintf("to have perimeter of %.5f", want),
			Matches:     got.Perimeter() == want,
			But:         fmt.Sprintf("it had perimeter of %.5f", got.Perimeter()),
			SubjectName: "Shape",
		}
	}
}

// https://stackoverflow.com/a/47969546
const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
