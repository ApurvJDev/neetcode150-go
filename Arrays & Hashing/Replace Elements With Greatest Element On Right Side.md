Replace Elements With Greatest Element On Right Side
Easy
Topics
Company Tags
You are given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.

Example 1:
Input: arr = [2,4,5,3,1,2]
Output: [5,5,3,2,2,-1]
Example 2:
Input: arr = [3,3]
Output: [3,-1]
Constraints:
1 <= arr.length <= 10,000
1 <= arr[i] <= 100,000

# Go
* [Go code](./code/replace_elements.go)
### O (n ^ 2)
```go
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	
	return b
}

func replaceElements(arr []int) []int {
	size := len(arr)
	for left := 0; left < size - 1; left++ {
		greatestNum := math.MinInt
		for right := left + 1; right < size; right++  {
			greatestNum = maxInt(arr[right], greatestNum)
		}
		arr[left] = greatestNum
		greatestNum = math.MinInt
	}
	arr[size - 1] = -1

	return arr
}
```
### O(n):
```go
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
```

# Cpp
* [Cpp code](./code/replace_elements.cpp)
```cpp
class Solution {
public:
    vector<int> replaceElements(vector<int>& arr) {
        int n = arr.size();
        vector<int> res(n);
        int maxNum = arr[n - 1];

        for(int i = n - 2; i >= 0; i--) {
            maxNum = max(maxNum, arr[i + 1]);
            res[i] = maxNum;
        }
        res[n - 1] = -1;

        return res;
    }
};
```
