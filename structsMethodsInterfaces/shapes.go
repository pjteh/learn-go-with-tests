package structsmethodsinterfaces

type Rectangle struct {
	Width  float64
	Length float64
}

// Calculates perimeter of a rectangle of length l and width w
func Perimeter(r Rectangle) float64 {
	return 2*r.Length + 2*r.Width
}

// Calculates area of rectangle
func Area(r Rectangle) float64 {
	return r.Length * r.Width
}
