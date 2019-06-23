/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
package leetcode

func isMatch(s string, p string) bool {
	return matchWithBacktracking(s, p)
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
