package structsmethodsinterfaces

// Calculates perimeter of a rectangle of length l and width w
func Perimeter(length, width float64) float64 {
	return 2*length + 2*width
}

// Calculates area of rectangle
func Area(length, width float64) float64 {
	return length * width
}
