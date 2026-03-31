package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("Repeat 5 times.", func(t *testing.T) {
		want := "aaaaa"
		got := Repeat("a", 5)

		if want != got {
			t.Errorf("Wanted %q, got %q", want, got)
		}
	})

	t.Run("Repeat a 3 times.", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		if want != got {
			t.Errorf("Wanted %q, got %q", want, got)
		}
	})

	t.Run("Repeat b 3 times.", func(t *testing.T) {
		got := Repeat("b", 3)
		want := "bbb"

		if want != got {
			t.Errorf("Wanted %q, got %q", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	out := Repeat("b", 3)
	fmt.Println(out)
	//Output: bbb
}
