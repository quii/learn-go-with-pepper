package structs_methods_and_interfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type EquilateralTriangle struct {
	Width float64
}

func (e EquilateralTriangle) Area() float64 {
	return math.Sqrt(3) / 4 * (e.Width * e.Width)
}

func (e EquilateralTriangle) Perimeter() float64 {
	return 3 * e.Width
}
