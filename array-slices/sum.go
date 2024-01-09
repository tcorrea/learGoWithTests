package arrayslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// As mentioned, slices have a capacity. If you have a slice with a capacity of 2 and try to do
	// mySlice[10] = 1 you will get a runtime error.
	// However, you can use the append function which takes a slice and a new value, then returns a
	// new slice with all the items in it.

	// create a new slice of ints with the same length as numbersToSum
	// sums := make([]int, len(numbersToSum))
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
		// sums[index] = Sum(numbers)
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))

	}

	return sums
}
