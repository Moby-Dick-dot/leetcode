package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	//定义一个dummyhead，解决当链表只有一个节点的情况，删除它要找到其前一个节点
	dummyhead := new(ListNode)
	dummyhead.Next = head //定义完虚拟头节点要让它指向头节点
	fast := dummyhead
	slow := dummyhead
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyhead.Next
}
