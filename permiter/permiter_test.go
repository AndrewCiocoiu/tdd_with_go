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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if want != got {
			t.Errorf("Got %g, wanted %g.", got, want)
		}
	}

	t.Run("Gives correct area for rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Gives correct area for circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
