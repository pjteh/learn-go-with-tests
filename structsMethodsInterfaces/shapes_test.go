package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	t.Run("test rectangular area", func(t *testing.T) {
		got := Area(20.0, 40.0)
		want := 800.0

		if got != want {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	})
}
