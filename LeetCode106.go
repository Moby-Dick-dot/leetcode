package main

//
//func buildTree(inorder []int, postorder []int) *TreeNode {
//	if len(postorder) == 0 {
//		return nil
//	}
//	node := &TreeNode{Val: 0, Left: nil, Right: nil}
//
//	node.Val = postorder[len(postorder)-1]
//	pos := locate(inorder, node.Val)
//
//	inorder_left := inorder[:pos]
//	inorder_right := inorder[pos+1:]
//
//	postorder_left := postorder[:len(inorder_left)]
//	postorder_right := postorder[len(inorder_left) : len(inorder_left)+len(inorder_right)]
//
//	node.Left = buildTree(inorder_left, postorder_left)
//	node.Right = buildTree(inorder_right, postorder_right)
//	return node
//}
//
//func locate(inorder []int, val int) int {
//	for i := 0; i < len(inorder); i++ {
//		if val == inorder[i] {
//			return i
//		}
//	}
//	return 0
//}
