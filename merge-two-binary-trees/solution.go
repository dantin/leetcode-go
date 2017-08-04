package p617

import (
	"github.com/dantin/leetcode-go/stack"
	"github.com/dantin/leetcode-go/tree"
)

func MergeTrees(t1 *tree.TreeNode, t2 *tree.TreeNode) *tree.TreeNode {
	if t1 == nil {
		return t2
	}

	ss := stack.New()
	ss.Push([]*tree.TreeNode{t1, t2})
	for !ss.IsEmpty() {
		t := ss.Pop().([]*tree.TreeNode)
		if t[0] == nil || t[1] == nil {
			continue
		}
		t[0].Val += t[1].Val
		if t[0].Left == nil {
			t[0].Left = t[1].Left
		} else {
			ss.Push([]*tree.TreeNode{t[0].Left, t[1].Left})
		}
		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			ss.Push([]*tree.TreeNode{t[0].Right, t[1].Right})
		}
	}
	return t1
}

// recursive version
func MergeTrees1(t1 *tree.TreeNode, t2 *tree.TreeNode) *tree.TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = MergeTrees(t1.Left, t2.Left)
	t1.Right = MergeTrees(t1.Right, t2.Right)
	return t1
}
