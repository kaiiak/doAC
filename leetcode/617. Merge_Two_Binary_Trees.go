// https://leetcode.com/problems/merge-two-binary-trees/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	root := new(TreeNode)
	if root1 != nil && root2 != nil {
		root.Val = root2.Val + root1.Val
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
        return root
	} else if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}
	return nil
}