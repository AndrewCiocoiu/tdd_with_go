package arrsum

func ArrSum(arr [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}

	return sum
}
