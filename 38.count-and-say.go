package leetcode

import "strconv"

func countAndSay(n int) string {
	s := []byte{'1'}
	for i := 1; i < n; i++ {
		s = say(s)
	}

	return string(s)
}

func say(s []byte) []byte {
	l := len(s)
	var d []byte
	var i, j int
	for j = 1; j < l; j++ {
		if s[j] != s[i] {
			d = strconv.AppendInt(d, int64(j-i), 10)
			d = append(d, s[i])
			i = j
		}
	}
	d = strconv.AppendInt(d, int64(j-i), 10)
	d = append(d, s[i])
	return d
}
