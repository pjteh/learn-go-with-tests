package arraysslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum 3 element array", func(t *testing.T) {
		array := []int{0, 10, 20}
		got := Sum(array)
		want := 30
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
	t.Run("sum 7 element array", func(t *testing.T) {
		array := []int{0, 1, 2, 3, 4, 5, 6}
		got := Sum(array)
		want := 21
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}
