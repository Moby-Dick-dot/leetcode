package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	//使用递归
	dummyhead := new(ListNode)
	dummyhead.Next = head
	cur := dummyhead

}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	
}
