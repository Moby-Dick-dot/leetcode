package main

import "math"

func findMode(root *TreeNode) []int {
	var prev *TreeNode
	count := 1
	maxCount := math.MinInt
	res := make([]int, 0)
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev == nil {
			count = 1
		} else if node.Val == prev.Val {
			count++
		} else {
			count = 1
		}
		if count >= maxCount {
			if count > maxCount && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			maxCount = count
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return res
}
