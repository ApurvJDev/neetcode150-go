func twoSum(nums []int, target int) []int {
    myMap := make(map[int]int)

	for index, value := range nums {
		num := target - value

		if i, ok := myMap[num]; ok {
			return []int{i, index}
		}

		myMap[value] = index
	}

	return nil
}
