// https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var sums []int64
	bfs(root, 0, &sums)

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})
	if k > len(sums) {
		return -1
	}
	return sums[k-1]
}

func bfs(root *TreeNode, level int, sum *[]int64) {
	if root == nil {
		return
	}
	if len(*sum) <= level {
		*sum = append(*sum, 0)
	}
	(*sum)[level] += int64(root.Val)
	bfs(root.Left, level+1, sum)
	bfs(root.Right, level+1, sum)
}