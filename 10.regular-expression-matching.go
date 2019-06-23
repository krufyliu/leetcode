/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
package leetcode

func isMatch10(s string, p string) bool {
	return matchDp(s, p)
}

func matchRecursion(s, p string) bool {
	ls, lp := len(s), len(p)

	if lp == 0 {
		return ls == 0
	}

	firstMatch := (ls != 0) && (s[0] == p[0] || p[0] == '.')

	// 如果存在'*'
	if lp > 1 && p[1] == '*' {
		return matchRecursion(s, p[2:]) || (firstMatch && matchRecursion(s[1:], p))
	}

	return firstMatch && matchRecursion(s[1:], p[1:])
}

// dp[i][j] represent whether s[i:], p[j:] is matched
func matchDp(s, p string) bool {
	ls, lp := len(s), len(p)

	dp := make([][]bool, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]bool, lp+1)
	}

	for i := ls; i >= 0; i-- {
		for j := lp; j >= 0; j-- {
			match := false
			if j == lp {
				match = i == ls
			} else {
				firstMatch := (i != ls) && (s[i] == p[j] || p[j] == '.')
				if j+1 < lp && p[j+1] == '*' {
					match = dp[i][j+2] || (firstMatch && dp[i+1][j])
				} else {
					match = firstMatch && dp[i+1][j+1]
				}
			}
			dp[i][j] = match
		}
	}

	return dp[0][0]
}
