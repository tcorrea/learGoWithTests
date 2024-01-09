package arrayslices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// array
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15
		got := Sum(numbers)

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		// slice
		numbers := []int{1, 2, 3}
		expected := 6
		got := Sum(numbers)

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})
}
