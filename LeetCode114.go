package main

func flatten(root *TreeNode) {
	var traversal func(node *TreeNode) *TreeNode
	traversal = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = traversal(node.Left)
		node.Right = traversal(node.Right)

		if node.Left == nil {
			return node
		}
		temp := node.Right
		node.Right = node.Left
		cur := node.Left
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = temp
		node.Left = nil

		return node
	}
	traversal(root)
}

//func flatten(root *TreeNode) {
//	var traversal func(node *TreeNode) *TreeNode
//	traversal = func(node *TreeNode) *TreeNode {
//
//		if node == nil {
//			return nil
//		}
//
//		left := traversal(node.Left)
//		right := traversal(node.Right)
//
//		if left == nil && right == nil {
//			return node
//		}
//		if left != nil {
//			node.Left = nil
//			node.Right = left
//			left.Right = right
//
//		} else {
//			node.Left = nil
//			node.Right = right
//		}
//
//		return node
//	}
//	root = traversal(root)
//}
