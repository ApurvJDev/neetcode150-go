Concatenation of Array
Easy
Topics
Company Tags
You are given an integer array nums of length n. Create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

Specifically, ans is the concatenation of two nums arrays.

Return the array ans.

Example 1:

Input: nums = [1,4,1,2]

Output: [1,4,1,2,1,4,1,2]
Example 2:

Input: nums = [22,21,20,1]

Output: [22,21,20,1,22,21,20,1]
# Go
* [Go code](./code/get_concatenation.go)
```go
func getConcatenation(nums []int) []int {
    size := len(nums)
	result := make([]int, 2*size)

	for index, value := range nums {
		result[index] = value
		result[index + size] = value 
	}

	return result
}
```

# Cpp
* [Cpp code](./code/get_concatenation.cpp)
```cpp
class Solution {
public:
    vector<int> getConcatenation(vector<int>& nums) {
        int n = nums.size();
        vector<int> res(2 * n);

        for (int i = 0; i < n; i++) {
            res[i] = nums[i];
            res[i + n] = nums[i];
        }

        return res;
    }
};
```
