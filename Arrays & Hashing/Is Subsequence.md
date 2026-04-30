Easy
You are given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
Input: s = "node", t = "neetcode"
Output: true
Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false

* [Go code](./code/is_subsequence.go)
# Go
O(n^2)
```go
func isSubsequence(s string, t string) bool {
	left := 0
	for _, value := range s {
		found := false
		for i := left; i < len(t); i++ {
			if value == rune(t[i]) {
				left = i + 1
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
```
O(n) 2 pointer
```go
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if (s[i] == t[j]) {
			i = i + 1
		}
		j = j + 1
	}
	return i == len(s)
}
```
