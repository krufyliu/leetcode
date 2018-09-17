package divide_two_integers

import (
	"testing"
)

func Test_divide(t *testing.T) {
	// log.Print(10 ^ 3)
	// log.Print(divide(1, 1))
	// return
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
		{
			name: "case 3",
			args: args{
				dividend: -7,
				divisor:  3,
			},
			want: -3,
		},
		{
			name: "case 4",
			args: args{
				dividend: -7,
				divisor:  -3,
			},
			want: 2,
		},
		{
			name: "case 5",
			args: args{
				dividend: intMin,
				divisor:  -1,
			},
			want: 2147483647,
		},
		{
			name: "case 6",
			args: args{
				dividend: -1,
				divisor:  -100,
			},
			want: 1,
		},
		{
			name: "case 7",
			args: args{
				dividend: 100,
				divisor:  -1000,
			},
			want: 0,
		},
		{
			name: "case 8",
			args: args{
				dividend: 1,
				divisor:  2,
			},
			want: 0,
		},
		{
			name: "case 9",
			args: args{
				dividend: 1,
				divisor:  1,
			},
			want: 1,
		},
		{
			name: "case 10",
			args: args{
				dividend: -3,
				divisor:  -1,
			},
			want: 3,
		},
		{
			name: "case 11",
			args: args{
				dividend: -2147483648,
				divisor:  -3,
			},
			want: 715827882,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
