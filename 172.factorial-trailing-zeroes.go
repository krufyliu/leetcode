/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
package leetcode

func trailingZeroes(n int) int {
	c := 0
	for n >= 5 {
		n = n / 5
		c += n
	}

	return c
}

func countN(n, m int) int {
	c := 0
	for n%m == 0 {
		c++
		n = n / m
	}
	return c
}
