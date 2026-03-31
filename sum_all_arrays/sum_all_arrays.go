package sumallarrays

// Will return a slice with each element being the sum of a slice that was provided
func SumAllArrays(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, ArrSum(numbers))
	}

	return sums
}

// Will return a slice with each element being the sum of a slice that was provided except the first element
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, ArrSum(tail))

	}

	return sums
}

func ArrSum(arr []int) int {
	sum := 0
	for _, nr := range arr {
		sum += nr
	}

	return sum
}
