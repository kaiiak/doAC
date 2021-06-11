// https://leetcode.com/problems/merge-two-sorted-lists/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	l := &ListNode{}
	if l1.Val < l2.Val {
		l.Val = l1.Val
		l.Next = mergeTwoLists(l1.Next, l2)
	} else {
		l.Val = l2.Val
		l.Next = mergeTwoLists(l1, l2.Next)
	}
	return l
}

// 循环
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l    = new(ListNode)
		root = l
	)
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			root.Next = l1
			l1 = l1.Next
		} else {
			root.Next = l2
			l2 = l2.Next
		}
		root = root.Next
	}
	if l1 == nil {
		root.Next = l2
	}
	if l2 == nil {
		root.Next = l1
	}
	return l.Next
}