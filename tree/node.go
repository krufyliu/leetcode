package tree

import (
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewFromString(s string) *TreeNode {
	s = strings.Replace(s, " ", "", -1)
	nodes := strings.Split(s, ",")
	if 
}
