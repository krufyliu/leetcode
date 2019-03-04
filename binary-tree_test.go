package leetcode

import (
	"testing"
)

func TestNewFromString(t *testing.T) {
	// type args struct {
	// 	s string
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
	// 		if got := NewFromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("NewFromString() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	root := NewBinaryTreeFromString("-10,-3,0,5,9")
	t.Error(root.String())
}
