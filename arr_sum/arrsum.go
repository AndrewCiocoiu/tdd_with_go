package arrsum

func ArrSum(arr []int) int {
	sum := 0
	for _, nr := range arr {
		sum += nr
	}

	return sum
}
