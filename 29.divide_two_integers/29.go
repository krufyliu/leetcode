package divide_two_integers

const intMax int = 1<<31 - 1
const intMin int = -((1 << 31) - 1) - 1

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == intMin && divisor == -1 {
		return intMax
	}

	division := 0
	flag := (dividend ^ divisor) > 0
	if (dividend > 0) || (dividend < 0 && divisor < 0) {
		absDividend, absDivisor := abs(dividend), abs(divisor)
		for absDividend >= absDivisor {
			division++
			absDividend -= absDivisor
		}
	} else {
		absDividend, absDivisor := abs(dividend), abs(divisor)
		for absDividend > 0 {
			division++
			absDividend -= absDivisor
		}
	}
	if flag {
		return division
	}
	return -division
}
