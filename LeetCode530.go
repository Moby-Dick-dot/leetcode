package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	var travel func(node *TreeNode)
	min := math.MaxInt
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}

		travel(node.Left)
		if prev != nil && node.Val-prev.Val < min { //二叉搜索树是有序的，前一个节点的值必然小于当前节点的值
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}

//func getMinimumDifference(root *TreeNode) int {
//	res := make([]int, 0)
//
//	inordecr(root, &res)
//
//	return getMinAbsSub(res)
//}
//
//func getMinAbsSub(res []int) int {
//	i := 0
//	j := 1
//	min := math.MaxInt
//	for j < len(res) {
//		if int(math.Abs(float64(res[i]-res[j]))) < min {
//			min = int(math.Abs(float64(res[i] - res[j])))
//		}
//		i++
//		j++
//	}
//	return min
//}
//
//func inordecr(node *TreeNode, res *[]int) {
//	if node == nil {
//		return
//	}
//
//	inordecr(node.Left, res)
//	*res = append((*res), node.Val)
//	inordecr(node.Right, res)
//}
