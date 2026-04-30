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
* [Go code](./code/valid_anagram.go)
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
*   [Cpp code](./code/valid_anagram.cpp)

```cpp
class Solution {
public:
    bool isAnagram(string s, string t) {
        unordered_map<char, int> sMap;
        unordered_map<char, int> tMap;

        for(char c : s) {
            sMap[c] += 1;
        }
        for(char c : t) {
            tMap[c] += 1;
        }
        
        if(sMap.size() != tMap.size()) {
            return false;
        }

        for(const auto& [key, value] : sMap) {
            if (tMap.contains(key)) {
                if(tMap[key] != sMap[key]) {
                    return false;
                }
            }
            else {
                return false;
            }
        }
        return true;
    }
};
```
