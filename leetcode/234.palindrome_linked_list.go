// https://leetcode.com/problems/palindrome-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
	// 栈方法
	if head == nil {
		return false
	}
	stacks := []int{}
	for head != nil {
		stacks = append(stacks, head.Val)
		head = head.Next
	}
	for i := 0; i < len(stacks)/2; i++ {
		if stacks[i] != stacks[len(stacks)-i-1] {
			return false
		}
	}
	return true
}