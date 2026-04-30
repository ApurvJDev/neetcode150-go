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
