/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name   string
		args   args
		wanted []int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 3},
			},
			wanted: []int{1, 3, 2},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 1, 5},
			},
			wanted: []int{1, 5, 1},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{3, 2, 1},
			},
			wanted: []int{1, 2, 3},
		},
		{
			name: "case 4",
			args: args{
				nums: []int{3},
			},
			wanted: []int{3},
		},
		{
			name: "case 5",
			args: args{
				nums: []int{1, 4, 5, 3, 2},
			},
			wanted: []int{1, 5, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.wanted) {
				t.Errorf("nextPermutation want %v, but got %v", tt.wanted, tt.args.nums)
			}
		})
	}
}
