package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				s: "    ",
			},
			want: 0,
		},
		{
			name: "case 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				s: "Hello ",
			},
			want: 5,
		},
		{
			name: "case 5",
			args: args{
				s: " Hello",
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				s: " Hello ",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
