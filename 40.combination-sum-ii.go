/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	lc := len(candidates)
	if lc == 0 {
		return nil
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var result [][]int
	var dfs func(index int, target int, curSet []int)
	dfs = func(index int, target int, curSet []int) {
		if target < 0 {
			return
		}

		if target == 0 && len(curSet) > 0 {
			result = append(result, curSet)
		}

		setLen := len(curSet)
		for i := index; i < lc; i++ {
			if candidates[i] > target {
				break
			}

			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			newSet := make([]int, setLen+1)
			copy(newSet[0:], curSet)
			newSet[setLen] = candidates[i]
			dfs(i+1, target-candidates[i], newSet)
		}
	}

	var set []int
	dfs(0, target, set)

	return result
}
