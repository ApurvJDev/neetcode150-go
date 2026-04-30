func appendCharacters(s string, t string) int {
    i, j := 0, 0
	for i < len(s) && j < len(t) {
		if (s[i] == t[j]) {
			j++
		}
		i++
	}
	if j == len(t) {
		return 0
	}
	if j == 0 {
		return len(t)
	}
	return len(t) - j
}
