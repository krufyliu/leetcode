/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */
package leetcode

func containsDuplicate(nums []int) bool {
    numsCount := map[int]int{}

    for _, v := range nums {
        numsCount[v]++
        if numsCount[v] > 1 {
            return true
        }
    }

    return false
}

