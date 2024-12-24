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
