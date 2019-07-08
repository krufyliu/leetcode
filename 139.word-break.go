/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */
package leetcode

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	for _, word := range wordDict {
		if strings.HasSuffix(s, word) {
			if len(s) == len(word) {
				return true
			}

			if wordBreak(s[0:len(s)-len(word)], wordDict) {
				return true
			}
		}
	}

	return false
}
