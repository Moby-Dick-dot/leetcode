package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//只能是指针相等，不能是值相等
	//定义两个指针，一个先走差的步数（你怎么知道差是多少,统计每个链表长度，然后用大的减小的，然后移动大的临时指针和小的临时指针到达同样的起点，然后一起向后移动）
	len_a := 0

	curA := headA

	for curA != nil {
		curA = curA.Next
		len_a++
	}

	len_b := 0
	curB := headB

	for curB != nil {
		curB = curB.Next
		len_b++
	}

	curA = headA
	curB = headB

	if len_a >= len_b {
		gap := len_a - len_b
		for i := 0; i < gap; i++ {
			curA = curA.Next
		}
	} else {
		gap := len_b - len_a
		for gap != 0 {
			curB = curB.Next
			gap--
		}
	}

	for curA != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}

	return nil
}
