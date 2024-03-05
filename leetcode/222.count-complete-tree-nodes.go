// https://leetcode.cn/problems/count-complete-tree-nodes/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) (total int) {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		total += countNodes(root.Left)
	}
	if root.Right != nil {
		total += countNodes(root.Right)
	}
	return total + 1
}