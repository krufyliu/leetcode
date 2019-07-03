/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
package leetcode

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				n: 3,
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				n: 5,
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				n: 25,
			},
			want: 6,
		},
		{
			name: "case 4",
			args: args{
				n: 1808548329,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
