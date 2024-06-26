package arraysandslices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func ExampleSum() {
	got := Sum([]int{1, 2})
	fmt.Println(got)
	// Output: 3
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !SlicesEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	got := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(got)
	// Output: [3 9]
}

func TestSlicesEqual(t *testing.T) {
	input := []int{3, 9, 5}
	expected := []int{3, 9, 5}

	got := SlicesEqual(input, expected)
	want := true

	if got != want {
		t.Errorf("got %t but want %t given %v wanted %v", got, want, input, expected)
	}
}

func ExampleSlicesEqual() {
	input := []int{3, 9, 5}
	expected := []int{3, 9, 5}

	areEqual := SlicesEqual(input, expected)
	fmt.Println(areEqual)
	// Output: true
}
