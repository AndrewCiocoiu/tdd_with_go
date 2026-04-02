package permiter

import "testing"

func TestPermiter(t *testing.T) {
	t.Run("Gives correct permiter", func(t *testing.T) {
		got := Permiter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("Wanted %.2f, got %.2f", want, got)
		}
	})
}
