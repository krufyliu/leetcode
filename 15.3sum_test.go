package leetcode

import (
	"testing"
)

func Test_threeSum(t *testing.T) {
	// type args struct {
	// 	nums []int
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want [][]int
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("threeSum() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	solution := threeSum([]int{-1, 0, 1, 2, -1, -4, -2})
	t.Errorf("len(solution) = %d", len(solution))
	for i := 0; i < len(solution); i++ {
		t.Errorf("%#v", solution[i])
	}
}
