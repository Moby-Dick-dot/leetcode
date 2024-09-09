package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	//处理空链表
	if head == nil {
		return true
	}

	//计算链表长度
	len_list := 0
	cur := head
	for cur != nil {
		len_list++
		cur = cur.Next
	}

	//重置cur指针到链表头
	cur = head
	st := []int{}

	//将前半部分入栈
	for i := 0; i < len_list/2; i++ {
		st = append(st, cur.Val)
		cur = cur.Next
	}

	//链表有奇数个节点，跳过中间结点
	if len_list%2 == 1 {
		cur = cur.Next
	}

	//遍历后半部分，与栈内元素比较，并弹出栈内元素
	for cur != nil {
		if len(st) == 0 {
			return false
		}
		if cur.Val != st[len(st)-1] {
			return false
		}
		st = st[:len(st)-1]
		cur = cur.Next
	}

	return len(st) == 0
}
