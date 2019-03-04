package validparentheses

func isValid(s string) bool {
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	l := len(s)

	if l == 0 {
		return true
	}

	stack := []byte{}

	var sp = -1

	for i := 0; i < l; i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
			sp++
		case ')', ']', '}':
			if sp < 0 {
				return false
			}

			if pairs[stack[sp]] != s[i] {
				return false
			}

			stack = stack[:sp]
			sp--

		}
	}

	if sp >= 0 {
		return false
	}

	return true
}
