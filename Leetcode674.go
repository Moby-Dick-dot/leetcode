package main

func findLengthOfLCIS(nums []int) int {
	//连续
	//dp[i] 表示以下标i为结尾的最长连续递增子序列的长度

	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	res := dp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
