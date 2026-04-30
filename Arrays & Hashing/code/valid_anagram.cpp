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
