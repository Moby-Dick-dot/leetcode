package main

//func rob(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	//dp[i]表示考虑下标为i（包括i）以内的房屋，最多可以偷窃的金额为dp[i]
//	dp := make([]int, len(nums))
//
//	dp[0] = nums[0]
//	dp[1] = max(nums[0], nums[1])
//
//	for i := 2; i < len(nums); i++ {
//		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
//	}
//
//	return dp[len(nums)-1]
//}
