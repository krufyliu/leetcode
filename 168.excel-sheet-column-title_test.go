/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

package leetcode

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		n: 1,
		// 	},
		// 	want: "A",
		// },
		// {
		// 	name: "case 2",
		// 	args: args{
		// 		n: 2,
		// 	},
		// 	want: "B",
		// },
		// {
		// 	name: "case 3",
		// 	args: args{
		// 		n: 27,
		// 	},
		// 	want: "AA",
		// },
		// {
		// 	name: "case 4",
		// 	args: args{
		// 		n: 28,
		// 	},
		// 	want: "AB",
		// },
		{
			name: "case 5",
			args: args{
				n: 703,
			},
			want: "AAA",
		},
		{
			name: "case 6",
			args: args{
				n: 52,
			},
			want: "AZ",
		},
		{
			name: "case 7",
			args: args{
				n: 26,
			},
			want: "Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
