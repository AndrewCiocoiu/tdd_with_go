package sumallarrays

import (
	"slices"
	"testing"
)

func TestSumAllArrays(t *testing.T) {
	t.Run("test 2 slices as input", func(t *testing.T) {
		got := SumAllArrays([]int{1, 2}, []int{3, 5})
		want := []int{3, 8}

		if !slices.Equal(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum all tails of 3 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 5, 3}, []int{9, 2, 1}, []int{14, 22, 55})
		want := []int{8, 3, 77}

		if !slices.Equal(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("SumAllTails with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{9, 2, 1})
		want := []int{0, 3}

		if !slices.Equal(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}

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
