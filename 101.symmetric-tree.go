package leetcode

import "container/list"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lq, rq := list.New(), list.New()
	lq.PushBack(root.Left)
	rq.PushBack(root.Right)

	for lq.Len() > 0 && rq.Len() > 0 {
		le, re := lq.Front(), rq.Front()
		lq.Remove(le)
		rq.Remove(re)
		li, ri := le.Value.(*TreeNode), re.Value.(*TreeNode)
		if li == nil && ri == nil {
			continue
		}

		if li != nil && ri != nil {
			if li.Val != ri.Val {
				return false
			}
			lq.PushBack(li.Left)
			lq.PushBack(li.Right)
			rq.PushBack(ri.Right)
			rq.PushBack(ri.Left)
		} else {
			return false
		}
	}

	return true
}

func isSymmetricRecursively(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricLR(root.Left, root.Right)
}

func isSymmetricLR(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l != nil && r != nil {
		if l.Val != r.Val {
			return false
		}
		return isSymmetricLR(l.Left, r.Right) && isSymmetricLR(l.Right, r.Left)
	}

	return false
}
