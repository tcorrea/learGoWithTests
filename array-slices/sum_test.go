package arrayslices

import (
	"reflect"
	"slices"
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

func TestSumAll(t *testing.T) {
	t.Run("sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		// expected := "manoel"

		if !slices.Equal(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

		// reflect.DeepEqual nao quebra comparando slice com string
		// usar slices.Equal
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})
}

func TestAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}
