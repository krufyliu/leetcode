/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (24.88%)
 * Total Accepted:    178.6K
 * Total Submissions: 711.6K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 *
 * Example 1:
 *
 *
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 *
 *
 * Example 2:
 *
 *
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 *
 *
 */
package leetcode

func longestValidParentheses(s string) int {
	l := len(s)
	longest := 0

	dp := make([]int, l)
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else if i-1 >= 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = 2 + dp[i-1]
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
			if dp[i] > longest {
				longest = dp[i]
			}
		}
	}
	return longest
}
