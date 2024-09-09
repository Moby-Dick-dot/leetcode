//package main
//
//func wiggleMaxLength(nums []int) int {
//	if len(nums) == 1 {
//		return 1
//	}
//
//	prediff := 0
//	curdiff := 0
//	result := 1
//	for i := 0; i < len(nums)-1; i++ {
//		curdiff = nums[i+1] - nums[i]                                       //更新curdiff
//		if (prediff >= 0 && curdiff < 0) || (prediff <= 0 && curdiff > 0) { //出现摆动,说明当前节点在摆动子序列
//			result++
//			prediff = curdiff //出现摆动prediff才进行更新
//		}
//	}
//
//	return result
//}
