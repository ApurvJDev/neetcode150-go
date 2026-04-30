func getConcatenation(nums []int) []int {
    size := len(nums)
	result := make([]int, 2*size)

	for index, value := range nums {
		result[index] = value
		result[index + size] = value 
	}

	return result
}
