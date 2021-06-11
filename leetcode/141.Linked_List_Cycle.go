// https://leetcode.com/problems/linked-list-cycle/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    var fast, slow = head, head.Next
    for fast != nil && slow != nil {
        if fast == slow {
            return true
        }
        fast = fast.Next
        slow = slow.Next
        if slow != nil {
            slow = slow.Next
        } else {
            break
        }
    }
    return false
}