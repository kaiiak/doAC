// https://leetcode.com/problems/intersection-of-two-linked-lists/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// O(n)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := map[*ListNode]struct{}{}
	for headA != nil {
		nodeMap[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := nodeMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    p1, p2 := headA, headB
    for {
        if p1 == p2 {
            return p1
        }
        if p1 == nil {
            p1 = headB
        } else {
            p1 = p1.Next
        }
        if p2 == nil {
            p2 = headA
        } else {
            p2 = p2.Next
        }
    }
    return nil
}