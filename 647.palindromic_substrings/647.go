package palindromic_substrings

func expand(s string, i, j int) int {
	var count = 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		count++
		i--
		j++
	}
	return count
}

func CountSubstrings(s string) int {
	var count = 0
	for i := 0; i < len(s); i++ {
		count += expand(s, i, i)
		count += expand(s, i, i+1)
	}
	return count
}
