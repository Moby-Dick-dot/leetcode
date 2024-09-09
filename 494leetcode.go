package main

import "math"

func findTargetSumWays(nums []int, target int) int {
	//向数组中的每个整数前面添加 加号 或者 减号， 使得结果等于 target，请问有多少种方法
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if int(math.Abs(float64(target))) > sum {
		return 0
	}

	if (sum+target)%2 != 0 {
		return 0
	}

	bagsize := (sum + target) / 2

	dp := make([]int, bagsize+1)

	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := bagsize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bagsize]
}
