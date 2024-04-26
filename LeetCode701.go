package main

//func insertIntoBST(root *TreeNode, val int) *TreeNode {
//	prev := &TreeNode{Val: 0, Left: nil, Right: nil}
//	var traversal func(cur *TreeNode, val int)
//	traversal = func(cur *TreeNode, val int) {
//		if cur == nil {
//			node := new(TreeNode)
//			node.Val = val
//			if prev.Val < val {
//				prev.Right = node
//			}
//			if val < prev.Val {
//				prev.Left = node
//			}
//			return
//		}
//
//		prev = cur
//		if cur.Val > val {
//			traversal(cur.Left, val)
//		}
//		if cur.Val < val {
//			traversal(cur.Right, val)
//		}
//		return
//	}
//	if root == nil {
//		root = &TreeNode{Val: val, Left: nil, Right: nil}
//	}
//	traversal(root, val)
//	return root
//}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode{Val: val, Left: nil, Right: nil}
		return node
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

//func insertIntoBST(root *TreeNode, val int) *TreeNode {
//	traversal(root,val)
//	return root
//}
//
//func traversal(cur *TreeNode, val int) {
//	if cur == nil {
//		return
//	}
//	if
//
//	if cur.Val > val{
//		traversal(cur.Left,val)
//	}
//	prev = cur
//
//	if cur.Val < val{
//		traversal(cur.Right,val)
//	}
//}
