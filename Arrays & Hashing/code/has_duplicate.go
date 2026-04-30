func hasDuplicate(nums []int) bool {
    myMap := make(map[int]bool)

    for _, value := range nums {
        if myMap[value] {
            return true // duplicate found
        } else {
            myMap[value] = true
        }
    }

    return false
}
