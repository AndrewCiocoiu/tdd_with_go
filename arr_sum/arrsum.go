package arrsum

func ArrSum(arr [5]int) int {
	sum := 0
	for _, nr := range arr {
		sum += nr
	}

	return sum
}
