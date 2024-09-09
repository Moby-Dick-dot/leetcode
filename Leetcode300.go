package main

func lengthOfLIS(nums []int) int {
	//dp[i] 下标从0到i的序列中以nums[i]最长递增子序列的长度
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	res := dp[0]

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res

}
