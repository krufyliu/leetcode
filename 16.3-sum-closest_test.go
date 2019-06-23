/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 0",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "case 1",
			args: args{
				nums:   []int{-1, 8, 7, 6, -1, 1},
				target: 1,
			},
			want: -1,
		},
		{
			name: "case 2",
			args: args{
				nums:   []int{-2, -3, 1, 5, -4, -6},
				target: -10,
			},
			want: -9,
		},
		{
			name: "case 3",
			args: args{
				nums:   []int{-1, -1, -1, -1, -1, -1},
				target: -10,
			},
			want: -3,
		},
		{
			name: "case 4",
			args: args{
				nums:   []int{1, 1, 1, -1, -1, -1},
				target: 10,
			},
			want: 3,
		},
		{
			name: "case 5",
			args: args{
				nums:   []int{1, 1, 1, -1, -1, -1},
				target: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
