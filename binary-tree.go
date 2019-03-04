package leetcode

/**
 * Definition for a binary tree node.
 */

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTreeFromString(s string) *TreeNode {
	if s == "" {
		return nil
	}

	strs := strings.Split(s, ",")

	q := list.New()
	v, err := strconv.Atoi(strs[0])
	if err != nil {
		panic(err)
	}

	root := &TreeNode{Val: v}
	q.PushBack(root)
	i := 1
	l := len(strs)

	for i < l {
		e := q.Front()
		q.Remove(e)
		cur := e.Value.(*TreeNode)
		if strs[i] != "null" {
			lv, err := strconv.Atoi(strs[i])
			if err != nil {
				panic(err)
			}

			cur.Left = &TreeNode{Val: lv}
			q.PushBack(cur.Left)
		}
		i++
		if i >= l {
			break
		}

		if strs[i] != "null" {
			rv, err := strconv.Atoi(strs[i])
			if err != nil {
				panic(err)
			}

			cur.Right = &TreeNode{Val: rv}
			q.PushBack(cur.Right)
		}
		i++
	}

	return root
}

func (root *TreeNode) String() string {
	if root == nil {
		return ""
	}

	q := list.New()
	q.PushBack(root)

	strs := []string{}
	for q.Len() > 0 {
		var curE = q.Front()
		q.Remove(curE)
		node := curE.Value.(*TreeNode)
		if node != nil {
			strs = append(strs, strconv.Itoa(node.Val))
			q.PushBack(node.Left)
			q.PushBack(node.Right)
		} else {
			strs = append(strs, "null")
		}
	}
	// trim tail null
	var i = len(strs) - 1
	for i >= 0 && strs[i] == "null" {
		i--
	}
	strs = strs[:i+1]

	return strings.Join(strs, ",")
}
