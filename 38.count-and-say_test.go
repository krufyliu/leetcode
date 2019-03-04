package leetcode

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				n: 1,
			},
			want: "1",
		},
		{
			name: "case 2",
			args: args{
				n: 2,
			},
			want: "11",
		},
		{
			name: "case 3",
			args: args{
				n: 3,
			},
			want: "21",
		},
		{
			name: "case 4",
			args: args{
				n: 4,
			},
			want: "1211",
		},
		{
			name: "case 5",
			args: args{
				n: 5,
			},
			want: "111221",
		},
		{
			name: "case 6",
			args: args{
				n: 6,
			},
			want: "312211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
