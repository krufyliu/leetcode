/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
package leetcode

func trailingZeroes(n int) int {
	// twoCount := 0
	fiveCount := 0

	prod := 5
	i := 1

	for prod <= n {
		if i%5 == 0 {
			fiveCount += countN(i, 5)
		}
		i++
		fiveCount++
		prod = 5 * i
	}
	return fiveCount
}

func countN(n, m int) int {
	c := 0
	for n%m == 0 {
		c++
		n = n / m
	}
	return c
}
