package main

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return res
	}
	//使用队列,用mod来时实现弹出队头对元素
	queue := []*TreeNode{}

	//先将gen元素放入队列中
	queue = append(queue, root)

	for len(queue) != 0 {
		length := len(queue)
		tempArr := []int{}

		for i := 0; i < length; i++ {
			//取队头元素
			node := queue[0]
			//弹出队头元素
			queue = queue[1:]

			tempArr = append(tempArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, tempArr)
	}

	return res
}
