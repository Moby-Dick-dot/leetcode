package main

//
//func constructMaximumBinaryTree(nums []int) *TreeNode {
//	if len(nums) == 0 {
//		return nil
//	}
//
//	pos := locate_max(nums)
//
//	node := &TreeNode{Val: 0, Left: nil, Right: nil}
//	node.Val = nums[pos]
//
//	nums_left := nums[:pos]
//	nums_right := nums[pos+1:]
//
//	node.Left = constructMaximumBinaryTree(nums_left)
//	node.Right = constructMaximumBinaryTree(nums_right)
//	return node
//}
//
//func locate_max(nums []int) int {
//	max_index := 0
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > nums[max_index] {
//			max_index = i
//		}
//	}
//
//	return max_index
//}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTree1(nums, 0, len(nums))
}

func constructMaximumBinaryTree1(nums []int, leftIndex int, rightIndex int) *TreeNode {
	//类似用数组构造二叉树的题目，每次分隔尽量不要定义新的数组，而是通过下标索引直接在原数组上操作，这样可以节约时间和空间上的开销。
}
