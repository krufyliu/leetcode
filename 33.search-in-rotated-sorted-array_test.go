/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
package leetcode

import "testing"

func Test_search(t *testing.T) {
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
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "case 1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "case 2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 6,
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 4,
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 2,
			},
			want: 6,
		},
		{
			name: "case 5",
			args: args{
				nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
				target: 8,
			},
			want: 4,
		},
		{
			name: "case 6",
			args: args{
				nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
				target: 1,
			},
			want: 5,
		},
		{
			name: "case 7",
			args: args{
				nums:   []int{4, 5, 6, 7, 8, 1, 2, 3},
				target: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
