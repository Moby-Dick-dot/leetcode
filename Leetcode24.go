package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyhead := new(ListNode)
	dummyhead.Next = head
	cur := dummyhead

	//每次只需要更新cur就行，因为是要交换cur后面两个节点的位置（可以通过cur得到），如果这两个节点有一个为空，那么交换就结束了
	for cur.Next != nil && cur.Next.Next != nil {
		temp1 := cur.Next           //提前保存指向节点1的指针，防止找不到
		temp3 := cur.Next.Next.Next //提前保存指向节点3的指针，防止找不到

		cur.Next = cur.Next.Next
		cur.Next.Next = temp1
		temp1.Next = temp3

		cur = cur.Next.Next
	}

	return dummyhead.Next
}
