package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("Repeat a 5 times", func(t *testing.T) {
		want := "aaaaa"
		got := Repeat("a")

		if want != got {
			t.Errorf("Wanted %q, got %q", want, got)
		}
	})
}
