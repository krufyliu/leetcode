package leetcode

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
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
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "case 2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "case 4",
			args: args{
				prices: []int{1, 4},
			},
			want: 3,
		},
		{
			name: "case 5",
			args: args{
				prices: []int{7, 6},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
