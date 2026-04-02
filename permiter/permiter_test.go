package permiter

import "testing"

func TestPermiter(t *testing.T) {
	t.Run("Gives correct permiter", func(t *testing.T) {
		got := Permiter(10.0, 10.0)
		want := 100.0

		if got != want {
			t.Errorf("Wanted %f, got %f", want, got)
		}
	})
}
