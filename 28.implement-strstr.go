package leetcode

func strStr(haystack string, needle string) int {
	return strStrSimple(haystack, needle)
}

func strStrSimple(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	lh, ln := len(haystack), len(needle)

	var i int
	for i = 0; i <= lh-ln; i++ {
		var j int
		for j = 0; j < ln; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == ln {
			return i
		}
	}

	return -1
}
