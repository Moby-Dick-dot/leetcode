package main

func maxSubArray(nums []int) int {
	//dp[i] 包括下标i之前的最大连续子序列的最大和为dp[i]
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]

	res := dp[0]
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
