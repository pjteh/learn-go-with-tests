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
	t.Run("test rectangular area", func(t *testing.T) {
		rectangle := Rectangle{20.0, 40.0}
		got := rectangle.Area(rectangle)
		want := 800.0

		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})
	// Test for circular area
	t.Run("circular area", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area(circle)
		want := math.Pi * circle.Radius * circle.Radius

		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}

	})
}
