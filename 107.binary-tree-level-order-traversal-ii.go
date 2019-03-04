package leetcode

import "container/list"

type levelTreeNode struct {
	level int
	node  *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := list.New()
	q.PushBack(&levelTreeNode{level: 1, node: root})
	level := 0
	var traversal [][]int
	for q.Len() > 0 {
		curE := q.Front()
		q.Remove(curE)
		node := curE.Value.(*levelTreeNode)
		if node.level == level {
			traversal[level-1] = append(traversal[level-1], node.node.Val)
		} else {
			traversal = append(traversal, []int{node.node.Val})
			level = node.level
		}

		if node.node.Left != nil {
			q.PushBack(&levelTreeNode{level: node.level + 1, node: node.node.Left})
		}

		if node.node.Right != nil {
			q.PushBack(&levelTreeNode{level: node.level + 1, node: node.node.Right})
		}
	}
	// reverse traverse result
	l := len(traversal)
	i, j := 0, l-1
	for i < j {
		traversal[i], traversal[j] = traversal[j], traversal[i]
		i++
		j--
	}

	return traversal
}
