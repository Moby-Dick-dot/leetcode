package main

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBST1(root, math.MinInt64, math.MaxInt64)
}

func isValidBST1(node *TreeNode, minVal, maxVal int) bool {
	if node == nil {
		return true
	}

	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}

	return isValidBST1(node.Left, minVal, node.Val) && isValidBST1(node.Right, node.Val, maxVal)
}

//
//func isValidBST(root *TreeNode) bool {
//	return isValidBST1(root)
//}
//
//func isValidBST1(node *TreeNode) bool {
//	//后序遍历
//	if node == nil {
//		return true
//	}
//	if node.Left == nil && node.Right == nil {
//		return true
//	}
//	if node.Left != nil && node.Right == nil {
//		if node.Val <= node.Left.Val {
//			return false
//		}
//	}
//	if node.Left == nil && node.Right != nil {
//		if node.Val >= node.Right.Val {
//			return false
//		}
//	}
//
//	if node.Val <= node.Left.Val || node.Val >= node.Right.Val {
//		return false
//	}
//
//	if node.Left != nil {
//		leftResult := isValidBST1(node.Left)
//		if leftResult == false {
//			return leftResult
//		}
//	}
//	if node.Right != nil {
//		rightResult := isValidBST1(node.Right)
//		if rightResult == false {
//			return rightResult
//		}
//	}
//	return true
//}
