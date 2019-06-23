/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
package leetcode

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var result [][]int

	var set []int

	var f func(candidates []int, index int, target int, curSet []int)
	f = func(candidates []int, index int, target int, curSet []int) {
		llset := len(curSet)

		// 找到target
		if target == 0 && len(curSet) != 0 {
			result = append(result, curSet)
			return
		}

		newSet := make([]int, llset+1)
		copy(newSet[0:], curSet)
		for i := index; i < len(candidates); i++ {
			// 当天填充数比余下的target大，说明后续都无解了
			if candidates[i] > target {
				break
			}

			// 填充此数
			newSet[llset] = candidates[i]
			f(candidates, i, target-candidates[i], newSet)
		}
	}

	f(candidates, 0, target, set)

	return result
}
