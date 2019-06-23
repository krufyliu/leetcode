/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
package leetcode

import "testing"

func Test_isMatch10(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 0",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "case 1",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "case 5",
			args: args{
				s: "",
				p: "c*b*a*.*",
			},
			want: true,
		},
		{
			name: "case 6",
			args: args{
				s: "",
				p: "",
			},
			want: true,
		},
		{
			name: "case 7",
			args: args{
				s: "s",
				p: "",
			},
			want: false,
		},
		{
			name: "case 8",
			args: args{
				s: "a",
				p: "a*a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch10(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
