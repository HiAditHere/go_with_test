package arrayslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("5 elements", func(t *testing.T) {
		numbers := []int{5, 4, 3, 2, 1}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("wanted '%d', got '%d'", want, got)
		}
	})

	t.Run("3 elements", func(t *testing.T) {
		numbers := []int{3, 2, 1}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("wanted '%d', got '%d'", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checksums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", got, want)
		}
	}

	t.Run("Test 1", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 9})
		want := []int{2, 9}

		checksums(t, got, want)
	})

	t.Run("Test 2", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checksums(t, got, want)
	})
}
