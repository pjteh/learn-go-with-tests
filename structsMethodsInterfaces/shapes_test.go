package structsmethodsinterfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	hasArea := 40.0
	if got != hasArea {
		t.Errorf("hasArea %.2f, got %.2f", hasArea, got)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 20, Length: 40}, hasArea: 800.00},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: math.Pi * 10.0 * 10.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, area := range areaTests {
		got := area.shape.Area()
		if got != area.hasArea {
			t.Errorf("hasArea %g, got %g", area.hasArea, got)
		}
	}
}
