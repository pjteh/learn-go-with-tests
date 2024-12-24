package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
