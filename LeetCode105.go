package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: 0, Left: nil, Right: nil}
	node.Val = preorder[0]

	pos := locate(inorder, node.Val)

	if pos == -1 {
		return nil
	}
	inorder_left := inorder[:pos]
	inorder_right := inorder[pos+1:]

	preorder_left := preorder[1 : len(inorder_left)+1]
	preorder_right := preorder[len(inorder_left)+1:]

	node.Left = buildTree(preorder_left, inorder_left)
	node.Right = buildTree(preorder_right, inorder_right)

	return node
}

func locate(inorder []int, val int) int {
	for i := 0; i < len(inorder); i++ {
		if val == inorder[i] {
			return i
		}
	}
	return -1
}
