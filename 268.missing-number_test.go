/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */
package leetcode

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{0, 1, 2, 3},
			},
			want: 4,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{0},
			},
			want: 1,
		},
		{
			name: "case 5",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
