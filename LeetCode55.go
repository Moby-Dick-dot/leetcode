//package main
//
//func canJump(nums []int) bool {
//	if len(nums) == 1 {
//		return true
//	}
//
//	cover := 0
//	for i := 0; i <= cover; i++ {
//		cover = max(i+nums[i], cover) //i+nums[i]表示跳跃的范围
//		if cover >= len(nums)-1 {
//			return true
//		}
//	}
//	return false
//}
