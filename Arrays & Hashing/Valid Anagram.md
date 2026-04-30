Valid Anagram
Easy
Topics
Company Tags
Hints
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: s = "racecar", t = "carrace"

Output: true
Example 2:

Input: s = "jar", t = "jam"

Output: false
Constraints:

s and t consist of lowercase English letters.

# Go
```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
// range over a string returns rune
	ms := make(map[rune]int)
	mt := make(map[rune]int)

	for _, v := range s {
		ms[v]++
	}

	for _, v := range t {
		mt[v]++
	}

	for k, v := range ms {
		if count, ok := mt[k]; ok {
			if (count != v) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
```

# Cpp
*   [View C++ Solution](./code/contains_duplicate.cpp)
