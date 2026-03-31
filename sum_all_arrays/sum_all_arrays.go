package sumallarrays

// Will return a slice with each element being the sum of a slice that was provided
func SumAllArrays(numbersToSum ...[]int) []int {
	nrOfSums := len(numbersToSum)
	sums := make([]int, nrOfSums)

	for i, numbers := range numbersToSum {
		sums[i] = ArrSum(numbers)
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
