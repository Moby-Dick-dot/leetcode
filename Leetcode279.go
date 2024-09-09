package main

import "math"

func numSquares(n int) int {
	//完全背包
	//首先要先找到小于n的完全平方数
	//最少min
	//dp[j] 表示和为j（装满）的完全平方数的最少数量

	nums := []int{}

	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	dp := make([]int, n+1)

	//初始化(第一个为0，其他全部为最大值)  非0下标的dp[j]一定要初始为最大值，这样dp[j]在递推的时候才不会被初始值覆盖。
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt32
	}

	for i, l := 0, len(nums); i < l; i++ {
		for j := nums[i]; j <= n; j++ {
			if dp[j-nums[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-nums[i]]+1)
			}
		}
	}

	if dp[n] == math.MaxInt32 {
		return -1
	}

	return dp[n]
}
