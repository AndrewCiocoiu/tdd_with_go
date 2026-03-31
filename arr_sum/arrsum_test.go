package arrsum

import "testing"

func TestArrSum(t *testing.T) {
	t.Run("Returns correct sum of elements in array", func(t *testing.T) {
		nums := [5]int{1, 2, 3, 4, 5}

		got := ArrSum(nums)
		want := 15

		if want != got {
			t.Errorf("Wanted %d, got %d, given %v", want, got, nums)
		}

	})
}
