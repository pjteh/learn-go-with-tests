package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum 3 element array", func(t *testing.T) {
		array := []int{0, 10, 20}
		got := Sum(array)
		want := 30
		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum slice of 3 elements", func(t *testing.T) {
		got := SumAll([]int{1, 1, 1})
		want := []int{3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, but got %v", want, got)
		}
	})
}
