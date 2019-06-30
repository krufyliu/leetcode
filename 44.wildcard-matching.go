/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
package leetcode

func isMatch(s string, p string) bool {
	return greedyMatch(s, p)
}

func matchWithBacktracking(s string, p string) bool {
	ls, lp := len(s), len(p)

	if lp == 0 {
		return ls == 0
	}

	if p[lp-1] == '*' {
		// 去除连续的*
		i := lp - 1
		for i >= 0 && p[i] == '*' {
			i--
		}

		if i < 0 {
			// 全是*
			return true
		}

		if ls == 0 {
			return matchWithBacktracking(s, p[:i+1])
		}

		if p[i] == '?' {
			return matchWithBacktracking(s[:ls-1], p[:i]) || matchWithBacktracking(s[:ls-1], p[:i+2])
		}

		j := ls - 1
		for ; j >= 0; j-- {
			if s[j] == p[i] {
				break
			}
		}

		if j < 0 {
			return false
		}

		return matchWithBacktracking(s[:j], p[:i]) || matchWithBacktracking(s[:j], p[:i+2])

	}

	// p还未匹配完
	if ls == 0 {
		return false
	}

	// 最后一个字符不匹配
	if p[lp-1] != '?' && s[ls-1] != p[lp-1] {
		return false
	}

	return matchWithBacktracking(s[:ls-1], p[:lp-1])
}

func greedyMatch(s, p string) bool {
	ls, lp := len(s), len(p)

	var matchAt, starAt = 0, -1

	var i, j int
	for i < ls {
		if j < lp && p[j] == '*' {
			matchAt = i
			starAt = j
			j++ // *从0次开始匹配
		} else if j < lp && (p[j] == '?' || s[i] == p[j]) {
			i++
			j++
		} else if starAt > -1 {
			matchAt++ // * 多匹配一个字符
			i = matchAt
			j = starAt + 1
		} else {
			return false
		}
	}

	for j < lp {
		if p[j] != '*' {
			return false
		}
		j++
	}

	return true
}

func matchByDynamic(s, p string) bool {
	ls, lp := len(s), len(p)

	dp := make([][]bool, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]bool, lp+1)
	}

	dp[ls][lp] = true

	for i := ls; i >= 0; i-- {
		for j := lp; j >= 0; j-- {
			if i < ls && j < lp && (p[j] == '?' || s[i] == p[j]) {
				dp[i][j] = dp[i+1][j+1]
			} else if j < lp && p[j] == '*' {
				dp[i][j] = (i < ls && dp[i+1][j]) || dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
