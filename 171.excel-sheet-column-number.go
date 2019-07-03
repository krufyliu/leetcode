/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */
package leetcode

func titleToNumber(s string) int {
	ls := len(s)

	pow := 1
	var sum int
	for i := ls - 1; i >= 0; i-- {
		sum += int(s[i]-'A'+1) * pow
		pow *= 26
	}

	return sum
}
