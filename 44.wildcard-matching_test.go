/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
package leetcode

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		/*
			{
				name: "case 1",
				args: args{
					s: "aa",
					p: "a",
				},
				want: false,
			},
			{
				name: "case 2",
				args: args{
					s: "aa",
					p: "*",
				},
				want: true,
			},
			{
				name: "case 3",
				args: args{
					s: "cb",
					p: "*a",
				},
				want: false,
			},
			{
				name: "case 4",
				args: args{
					s: "adceb",
					p: "*a*b",
				},
				want: true,
			},
			{
				name: "case 5",
				args: args{
					s: "acdcb",
					p: "a*c?b",
				},
				want: false,
			},
			{
				name: "case 6",
				args: args{
					s: "",
					p: "*",
				},
				want: true,
			},
			{
				name: "case 7",
				args: args{
					s: "",
					p: "***",
				},
				want: true,
			},
			{
				name: "case 8",
				args: args{
					s: "a",
					p: "***?",
				},
				want: true,
			},
			{
				name: "case 9",
				args: args{
					s: "a",
					p: "**?**",
				},
				want: true,
			},
			{
				name: "case 10",
				args: args{
					s: "",
					p: "**?**",
				},
				want: false,
			},
			{
				name: "case 11",
				args: args{
					s: "aaabaababaafadfadbbaafadafadfadfbbbaaaaba",
					p: "a****************************b",
				},
				want: false,
			},
			{
				name: "case 12",
				args: args{
					s: "babbbbaabababaabbababaababaabbaabababbaaababbababaaaaaabbabaaaabababbabbababbbaaaababbbabbbbbbbbbbaabbb",
					p: "b**bb**a**bba*b**a*bbb**aba***babbb*aa****aabb*bbb***a*************",
				},
				want: false,
			},
			{
				name: "case 13",
				args: args{
					s: "aa",
					p: "a*",
				},
				want: true,
			},
			{
				name: "case 14",
				args: args{
					s: "aa",
					p: "b*",
				},
				want: false,
			},
			{
				name: "case 15",
				args: args{
					s: "aa",
					p: "*a",
				},
				want: true,
			},
			{
				name: "case 16",
				args: args{
					s: "aa",
					p: "*b",
				},
				want: false,
			},
			{
				name: "case 17",
				args: args{
					s: "mississippi",
					p: "m*issip*",
				},
				want: true,
			},
		*/
		{
			name: "case 18",
			args: args{
				s: "abaabaaaabbabbaaabaabababbaabaabbabaaaaabababbababaabbabaabbbbaabbbbbbbabaaabbaaaaabbaabbbaaaaabbbabb",
				p: "ab*aaba**abbaaaa**b*b****aa***a*b**ba*a**ba*baaa*b*ab*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
