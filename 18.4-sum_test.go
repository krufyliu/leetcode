/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (29.64%)
 * Total Accepted:    216.9K
 * Total Submissions: 725.7K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 *
 * Note:
 *
 * The solution set must not contain duplicate quadruplets.
 *
 * Example:
 *
 *
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 *
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		nums:   []int{1, 0, -1, 0, -2, 2},
		// 		target: 0,
		// 	},
		// 	want: [][]int{
		// 		[]int{-1, 0, 0, 1},
		// 		[]int{-2, -1, 1, 2},
		// 		[]int{-2, 0, 0, 2},
		// 	},
		// },
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		nums:   []int{-3, -1, 0, 2, 4, 5},
		// 		target: 0,
		// 	},
		// 	want: [][]int{
		// 		[]int{-3, -1, 0, 4},
		// 	},
		// },
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		nums:   []int{-1, -5, -5, -3, 2, 5, 0, 4},
		// 		target: -7,
		// 	},
		// 	want: [][]int{
		// 		[]int{-5, -5, -1, 4},
		// 		[]int{-5, -3, -1, 2},
		// 	},
		// },
		{
			name: "case 1",
			args: args{
				nums:   []int{0, 0, 0, 0},
				target: 0,
			},
			want: [][]int{
				[]int{0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
