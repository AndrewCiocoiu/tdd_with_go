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
	t.Run("Gives correct area for rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("Wanted %.2f, got %.2f", want, got)
		}
	})

	t.Run("Gives correct area for circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("Wanted %g, got %g", want, got)
		}
	})
}
