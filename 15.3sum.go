package leetcode

import "sort"

// a + b + c = 0
// a + b = -c
func threeSum(nums []int) [][]int {
	l := len(nums)

	solution := [][]int{}
	if l < 3 {
		return solution
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < l-2; i++ {
		j, k := i+1, l-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				solution = append(solution, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}

			}
		}

		for i < l-2 && nums[i] == nums[i+1] {
			i++
		}
	}

	return solution
}
