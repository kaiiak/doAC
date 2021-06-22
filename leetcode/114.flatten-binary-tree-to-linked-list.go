// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode) {
	// 先序遍历
	zip(root)
}

func zip(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := root.Right
	root.Right = zip(root.Left)
	root.Left = nil
	if right == nil {
		return root
	}
	var (
		temp = root.Right
		pre  = root
	)
	for temp != nil {
		pre = temp
		temp = temp.Right
	}
	pre.Right = zip(right)
	return root
}