package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {

	t.Run("2 + 2 should be 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("Expected %d got %d.", expected, sum)
		}
	})

	t.Run("8 + -10 should be -2", func(t *testing.T) {
		sum := Add(8, -10)
		expected := -2
		if sum != expected {
			t.Errorf("Expected %d got %d.", expected, sum)
		}
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
