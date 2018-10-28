package repeated_dns_sequences

func simpleHash(s string) int {
	var hash = 0
	for i := 0; i < len(s); i++ {
		hash += int(s[i])
	}
	return hash
}

func FindRepeatedDnaSequences(s string) []string {
	substrings := make([]string, 0, 4)
	var set = make(map[string]int)
	length := len(s)
	for i := 0; i+10 <= length; i++ {
		set[s[i:i+10]] = set[s[i:i+10]] + 1
	}

	for key, v := range set {
		if v > 1 {
			substrings = append(substrings, key)
		}
	}
	return substrings
}
