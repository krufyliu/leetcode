package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{}
	for i := 0; i < 1000; i++ {
		tests = append(tests, test{

			name: fmt.Sprintf("case %d", i),
			args: args{
				x: i,
			},
			want: int(math.Sqrt(float64(i))),
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
