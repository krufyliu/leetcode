/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */
/*
 * 余九定理
 * 树根 dr(n) = 1 + (n-1) % 9
 */
package leetcode

func addDigits(num int) int {
	return 1 + (num-1)%9
}
