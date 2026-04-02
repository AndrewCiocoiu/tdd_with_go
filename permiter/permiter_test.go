package permiter

import "testing"

func TestPermiter(t *testing.T) {
	t.Run("Gives correct permiter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Permiter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("Wanted %.2f, got %.2f", want, got)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("Gives correct area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		want := 100.0

		if got != want {
			t.Errorf("Wanted %.2f, got %.2f", want, got)
		}
	})
}
