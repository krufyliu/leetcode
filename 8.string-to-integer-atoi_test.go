/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
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
				str: "42",
			},
			want: 42,
		},
		{
			name: "case 1",
			args: args{
				str: "        -42",
			},
			want: -42,
		},
		{
			name: "case 2",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "case 3",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "case 5",
			args: args{
				str: "91283472332",
			},
			want: 2147483647,
		},
		{
			name: "case 6",
			args: args{
				str: "  ",
			},
			want: 0,
		},
		{
			name: "case 7",
			args: args{
				str: "adfadfad",
			},
			want: 0,
		},
		{
			name: "case 8",
			args: args{
				str: "        -wffwf",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
