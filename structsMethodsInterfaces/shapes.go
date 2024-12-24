package structsmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

// Calculates perimeter of a rectangle of length l and width w
func Perimeter(r Rectangle) float64 {
	return 2*r.Length + 2*r.Width
}

// Calculates area of rectangle
func (rect Rectangle) Area() float64 {
	return rect.Length * rect.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
