package sumallarrays

import (
	"slices"
	"testing"
)

func TestSumAllArrays(t *testing.T) {
	t.Run("Sum 2 arrays", func(t *testing.T) {
		got := SumAllArrays([]int{1, 2}, []int{3, 5})
		want := []int{4, 7}

		if !slices.Equal(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
