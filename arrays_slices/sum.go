package arraysslices

func Sum(arr [3]int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
