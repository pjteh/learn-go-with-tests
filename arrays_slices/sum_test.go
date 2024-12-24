package arraysslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum 3 element array", func(t *testing.T) {
		array := [3]int{0, 10, 20}
		got := Sum(array)
		want := 30
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}
