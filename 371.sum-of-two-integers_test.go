/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

package leetcode

import "testing"

func Test_getSum(t *testing.T) {
	var target int = (1 << 63) - 1
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				a: -2,
				b: 3,
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				a: -123123,
				b: 123123,
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				a: (1 << 63) - 1,
				b: 1,
			},
			want: target + 1,
		},
		{
			name: "case 4",
			args: args{
				a: (1 << 63) - 2,
				b: 1,
			},
			want: target,
		},
		{
			name: "case 5",
			args: args{
				a: -((1 << 63) - 1),
				b: -2,
			},
			want: -target - 2,
		},
		{
			name: "case 6",
			args: args{
				a: -1,
				b: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
