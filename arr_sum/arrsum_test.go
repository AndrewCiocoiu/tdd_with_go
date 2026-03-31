package arrsum

import "testing"

func TestArrSum(t *testing.T) {
	t.Run("correct sum of 5 elements", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := ArrSum(nums)
		want := 15

		if want != got {
			t.Errorf("Wanted %d, got %d, given %v", want, got, nums)
		}

	})

	t.Run("correct sum of arbitrary nr of elements", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := ArrSum(nums)
		want := 6

		if want != got {
			t.Errorf("Wanted %d, got %d, given %v", want, got, nums)
		}
	})
}
