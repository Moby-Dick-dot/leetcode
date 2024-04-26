package main

func searchBST(root *TreeNode, val int) *TreeNode {

	return searchBST1(root, val)
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil && root.Val == val {
		return root
	}

	result := new(TreeNode)
	if val > root.Val {
		result = searchBST1(root.Right, val)
	}
	if val < root.Val {
		result = searchBST1(root.Left, val)
	}
	return result
}

//func searchBST1(root *TreeNode, val int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	if root.Val == val {
//		return root
//	}
//
//	if root.Left != nil {
//		leftResult := searchBST1(root.Left, val)
//		if leftResult != nil {
//			return leftResult
//		}
//	}
//	if root.Right != nil {
//		rightResult := searchBST1(root.Right, val)
//		if rightResult != nil {
//			return rightResult
//		}
//	}
//	return nil
//}
