package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	l := len(strs)

	if l == 0 {
		return ""
	}

	if l == 1 {
		return strs[0]
	}

	comm := strs[0]
	for i := 1; i < l; i++ {
		comm = longestCommonPrefixOfTwo(comm, strs[i])
		if len(comm) == 0 {
			return comm
		}
	}

	return comm
}

func longestCommonPrefixOfTwo(a, b string) string {
	la := len(a)
	lb := len(b)

	var t []byte
	for i, j := 0, 0; i < la && j < lb; {
		if a[i] == b[j] {
			t = append(t, a[i])
		} else {
			break
		}
		i++
		j++
	}

	return string(t)
}
