package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		if got != want {
			t.Errorf("expect %v, but got %v", want, got)
		}
	})
	t.Run("repeat 1 times", func(t *testing.T) {
		got := Repeat("c", 1)
		want := "c"
		if got != want {
			t.Errorf("expect %v, but got %v", want, got)
		}
	})
}

func ExampleRepeat() {
	str := Repeat("ds", 4)
	fmt.Println(str)
	//Output: dsdsdsds
}
