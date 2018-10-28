package ZigZagConversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l := len(s)
	bs := make([]byte, l)
	cur := 0
	cycle := numRows + (numRows - 2)
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < l; {
				bs[cur] = s[j]
				cur++
				j += cycle
			}

		} else {
			step := cycle - i - i
			for j := i; j < l; {
				bs[cur] = s[j]
				cur++
				j += step
				step = cycle - step
			}
		}

	}

	return string(bs)
}
