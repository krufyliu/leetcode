/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */
package leetcode

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    3,
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    3,
			},
			want: 5,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{2, 1},
				k:    1,
			},
			want: 2,
		},
		{
			name: "case 6",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    1,
			},
			want: 6,
		},
		{
			name: "case 7",
			args: args{
				nums: []int{99, 99},
				k:    1,
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
