package p617

import "github.com/dantin/leetcode-go/tree"

func MergeTrees(t1 *tree.TreeNode, t2 *tree.TreeNode) *tree.TreeNode {
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
