package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SlicesEqual compares the values of each slices
// and returns 'false' if one of the values are not equal
func SlicesEqual(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for key := range got {
		if got[key] != want[key] {
			return false
		}
	}
	return true
}
