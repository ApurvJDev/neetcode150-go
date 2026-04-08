Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

Example 1:

Input: nums = [1, 2, 3, 3]

Output: true

Example 2:

Input: nums = [1, 2, 3, 4]

Output: false

map[int]int approach:
```go
func hasDuplicate(nums []int) bool {
    myMap := make(map[int]int)

    for index, value := range nums {
        if _, ok := myMap[value]; ok {
            return true // duplicate found
        } else {
            myMap[value] = index
        }
    }

    return false
}
```

map[int]bool approach (Cleaner approach):
```go
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
```
