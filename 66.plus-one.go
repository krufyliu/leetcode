package leetcode

func plusOne(digits []int) []int {
	l := len(digits)
	carry := 1
	next := make([]int, l+1)
	copy(next[1:], digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i]+carry == 10 {
			next[i+1] = 0
		} else {
			next[i+1] = digits[i] + carry
			carry = 0
			break
		}
	}

	if carry > 0 {
		next[0] = carry
		return next
	}

	return next[1:]
}
