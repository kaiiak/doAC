// https://leetcode.com/problems/diameter-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
        return 0
    }
    weight := treeWeight(root)
    left := treeWeight(root.Left)
    right := treeWeight(root.Right)
	return max(weight, left, right, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func treeWeight(root *TreeNode) int {
    if root == nil {
		return 0
	}
	var (
		left, right, weight = 0, 0, 0
	)
	left += treeDeep(root.Left)
	right += treeDeep(root.Right)
	weight = left + right
	// fmt.Println(left, right, weight)
	return weight
}

func treeDeep(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var (
        left = treeDeep(root.Left)
        right = treeDeep(root.Right)
    )
    return max(left,right)+1
}

func max(nums ...int) int {
    max := nums[0]
    for i := 0; i < len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}