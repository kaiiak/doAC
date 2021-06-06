/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSymmetrices(root.Left, root.Right)
}

func isSymmetrices(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if (left != nil && right == nil ) || (left == nil && right != nil ) {
        return false
    }
    if left.Val != right.Val {
        return false
    }
    return isSymmetrices(left.Left, right.Right) && isSymmetrices(left.Right , right.Left)
}