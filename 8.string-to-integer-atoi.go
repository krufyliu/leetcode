/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package leetcode

const INT_MAX = (1 << 31) - 1
const INT_MIN = -(1 << 31)

func myAtoi(str string) int {
	ll := len(str)
	var i int
	var sign = 1
	// find first non-whitesapce character
	for i = 0; i < ll; i++ {
		if str[i] != ' ' {
			break
		}
	}

	// can't not find non-whitespace character
	if i == ll {
		return 0
	}

	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}

	//
	if i == ll {
		return 0
	}

	var num int64
	for ; i < ll; i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			break
		}
		num = (num * 10) + int64(str[i]-'0')
		if sign == -1 && num > -INT_MIN {
			return INT_MIN
		} else if sign == 1 && num > INT_MAX {
			return INT_MAX
		}
	}

	return int(num) * sign
}
