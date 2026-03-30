func buildFailure(p string) []int {
	fail := make([]int, len(p))
	j := 0
	for i := 1; i < len(p); i++ {
		for j > 0 && p[i] != p[j] {
			j = fail[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		fail[i] = j
	}
	return fail
}

func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	fail := buildFailure(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = fail[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}