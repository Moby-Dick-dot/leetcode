package main

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > high {
		root = trimBST(root.Left, low, high)
	} else if root.Val < low {
		root = trimBST(root.Right, low, high)
	} else if root.Val == low {
		root.Right = trimBST(root.Right, low, high)
		root.Left = nil
	} else if root.Val == high {
		root.Left = trimBST(root.Left, low, high)
		root.Right = nil
	} else {
		root.Left = trimBST(root.Left, low, root.Val)
		root.Right = trimBST(root.Right, root.Val, high)
	}

	return root
}
