/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
package leetcode

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				word1: "abc",
				word2: "abc",
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				word1: "0",
				word2: "abc",
			},
			want: 3,
		},
		{
			name: "case 5",
			args: args{
				word1: "abcde",
				word2: "abcd",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
