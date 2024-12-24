package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}

const pi = 3.14159

// Calculates perimeter of a rectangle of length l and width w
func Perimeter(r Rectangle) float64 {
	return 2*r.Length + 2*r.Width
}

// Calculates area of rectangle
func (rect *Rectangle) Area(r Rectangle) float64 {
	return r.Length * r.Width
}

func (c *Circle) Area(circle Circle) float64 {
	return pi * c.Radius * c.Radius
}
