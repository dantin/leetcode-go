package p617

import (
	"testing"

	"github.com/dantin/leetcode-go/tree"
)

func TestMergeTrees(t *testing.T) {
	var left, right, node *tree.TreeNode

	left, right = &tree.TreeNode{Val: 1}, nil

	node = MergeTrees(left, right)
	if 1 != node.Val {
		t.Error("merge two binary trees failed")
	}

	left, right = &tree.TreeNode{Val: 1}, &tree.TreeNode{Val: 1}
	node = MergeTrees(left, right)
	if 2 != node.Val {
		t.Error("merge two binary trees failed")
	}
}
