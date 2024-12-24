package arraysslices

func Sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
func SumAll(intSlice ...[]int) []int {
	returnSlice := []int{}
	for _, value := range intSlice {
		sum := Sum(value)
		returnSlice = append(returnSlice, sum)
	}
	return returnSlice
}

func SumAllTails(intSlice ...[]int) []int {
	returnSlice := []int{}
	for _, value := range intSlice {
		tail := value[1:]
		returnSlice = append(returnSlice, Sum(tail))
	}
	return returnSlice
}
