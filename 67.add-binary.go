package leetcode

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)

	var sum []byte
	if la > lb {
		sum = make([]byte, la+1)
	} else {
		sum = make([]byte, lb+1)
	}

	var carry, cur byte
	i, j, k := la-1, lb-1, len(sum)-1

	for i >= 0 && j >= 0 {
		cur = (a[i] - '0') + (b[j] - '0') + carry
		if cur > 1 {
			carry = 1
		} else {
			carry = 0
		}
		sum[k] = (cur & 1) + '0'
		i--
		j--
		k--
	}

	if i >= 0 {
		for i >= 0 {
			cur = (a[i] - '0') + carry
			if cur > 1 {
				carry = 1
			} else {
				carry = 0
			}
			sum[k] = (cur & 1) + '0'
			k--
			i--
		}
	}

	if j >= 0 {
		for j >= 0 {
			cur = (b[j] - '0') + carry
			if cur > 1 {
				carry = 1
			} else {
				carry = 0
			}
			sum[k] = (cur & 1) + '0'
			k--
			j--
		}
	}

	if carry > 0 {
		sum[0] = '1'
		return string(sum)
	}

	return string(sum[1:])
}
