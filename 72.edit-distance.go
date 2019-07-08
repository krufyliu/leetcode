/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
package leetcode

func minDistance(word1 string, word2 string) int {
	lw1, lw2 := len(word1), len(word2)
	dp := make([][]int, lw1+1)
	for i := 0; i <= lw1; i++ {
		dp[i] = make([]int, lw2+1)
	}

	for i := 0; i <= lw1; i++ {
		for j := 0; j <= lw2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = dp[i-1][j-1]
					if dp[i-1][j] < dp[i][j] {
						dp[i][j] = dp[i-1][j]
					}
					if dp[i][j-1] < dp[i][j] {
						dp[i][j] = dp[i][j-1]
					}
					dp[i][j]++
				}
			}
		}
	}
	return dp[lw1][lw2]
}
