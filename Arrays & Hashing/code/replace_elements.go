func replaceElements(arr []int) []int {
	size := len(arr)
	res := make([]int, size)
	num := math.MinInt
	
	for right := size - 2; right >=0 ; right-- {
		num = maxNum(num, arr[right + 1])
		res[right] = num
	}

	res[size - 1] = -1

	return res
}

func maxNum (a, b int) int {
	if a > b {
		return a
	}
	return b
}
