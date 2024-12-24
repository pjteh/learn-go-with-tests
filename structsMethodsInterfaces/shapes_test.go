package structsmethodsinterfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{20, 40}, 800.00},
		{Circle{10}, math.Pi * 10.0 * 10.0},
		{Triangle{12, 6}, 36.0},
	}

	for _, area := range areaTests {
		got := area.shape.Area()
		if got != area.want {
			t.Errorf("want %g, got %g", area.want, got)
		}
	}
}
