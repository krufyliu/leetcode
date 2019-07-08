/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "case 1",
		// 	args: args{
		// 		root: NewBinaryTreeFromString("1,null,2,3"),
		// 	},
		// 	want: []int{3, 2, 1},
		// },
		{
			name: "case 2",
			args: args{
				root: NewBinaryTreeFromString("1,2,3,4,5,null,null,null,null,6,7"),
			},
			want: []int{4, 6, 7, 5, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
