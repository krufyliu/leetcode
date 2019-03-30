package leetcode

func isPalindrome(s string) bool {
	l := len(s)

	var i, j = 0, l - 1

	for i < j {
		for i < j && !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9')) {
			i++
		}
		for i < j && !((s[j] >= 'a' && s[j] <= 'z') || (s[j] >= 'A' && s[j] <= 'Z') || (s[j] >= '0' && s[j] <= '9')) {
			j--
		}
		if i < j {
			a, b := s[i], s[j]
			if a >= 'a' {
				a = a - 32
			}
			if b >= 'a' {
				b = b - 32
			}

			if a != b {
				return false
			}
		}
		i++
		j--
	}

	return true
}
