package main

//
//func rob(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	//那2次只截取数组的一部分（第一次不要最后一个元素，第二次不要第一个元素）
//	//这样就不会出现同时选中第一个和最后一个元素的情况
//
//	res1 := robhouse(nums, 0, len(nums)-2)
//	res2 := robhouse(nums, 1, len(nums)-1)
//	return max(res1, res2)
//}
//
//func robhouse(nums []int, start int, end int) (res int) {
//	if start == end {
//		return nums[start]
//	}
//
//	nums = nums[start : end+1]
//	//dp[i] 考虑下标在i（包括i）以内的房屋，所能偷窃的最大金额
//	dp := make([]int, len(nums))
//	dp[0] = nums[0]
//	dp[1] = max(nums[0], nums[1])
//
//	for i := 2; i < len(nums); i++ {
//		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
//	}
//
//	return dp[len(nums)-1]
//}
