// https://leetcode.cn/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	var nodes = map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := nodes[head]; ok {
			return head
		}
		nodes[head] = struct{}{}
		head = head.Next
	}
	return nil
}
// 还可以使用双指针方法，一个快指针，一个慢指针，快指针每次走两步，慢指针每次走一步，如果有环，两个指针一定会相遇。