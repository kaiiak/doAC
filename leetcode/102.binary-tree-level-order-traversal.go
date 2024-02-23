// https://leetcode.cn/problems/binary-tree-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) (nodes [][]int) {
	if root == nil {
		return
	}

	nodes = make([][]int, 0)

	level(0, root, &nodes)

	return
}

func level(l int, root *TreeNode, nodes *[][]int) {
	if root == nil {
		return
	}

	if len(*nodes) <= l {
		*nodes = append(*nodes, []int{})
	}

	(*nodes)[l] = append((*nodes)[l], root.Val)
	level(l+1, root.Left, nodes)
	level(l+1, root.Right, nodes)
}
