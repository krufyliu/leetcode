package leetcode

func lengthOfLastWord(s string) int {
	l := len(s)

	var c int
	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if c > 0 {
				break
			}
		} else {
			c++
		}
	}

	return c
}
