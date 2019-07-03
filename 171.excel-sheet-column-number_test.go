/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */
package leetcode

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				s: "A",
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				s: "Z",
			},
			want: 26,
		},
		{
			name: "case 3",
			args: args{
				s: "AA",
			},
			want: 27,
		},
		{
			name: "case 4",
			args: args{
				s: "AAA",
			},
			want: 703,
		},
		{
			name: "case 5",
			args: args{
				s: "AZ",
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
