package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				s: "",
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				s: "  , , , ,",
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				s: "ab2a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
