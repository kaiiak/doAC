// https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    var (
        cur = head
        next = head.Next
        first = true
    )
    if next == nil {
        return head
    }
    for {
        temp := next.Next
        next.Next = cur
        if first {
            cur.Next = nil
            first = false
        }
        cur = next
        next = temp
        if temp == nil {
            break
        }
    }
    return cur
}