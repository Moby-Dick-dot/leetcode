package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}
	//快慢指针
	fast := head.Next
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		//快指针一次走两步，慢指针一次走一步，如果有环，快指针肯定会追上满慢指针
		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
