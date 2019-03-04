package leetcode

import (
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	// type args struct {
	// 	nums []int
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want *TreeNode
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	t.Error(root.String())
}
