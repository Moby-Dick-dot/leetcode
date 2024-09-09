package main

func combinationSum4(nums []int, target int) int {
	//dp[j] 含义：当背包容量为j时，装满背包有dp[j]种方法（排列数）
	dp := make([]int, target+1)

	//初始化
	dp[0] = 1

	//遍历顺序：因为是求排列数，所以先遍历背包，后遍历物品
	for j := 0; j <= target; j++ { //遍历背包
		for i, l := 0, len(nums); i < l; i++ { //遍历物品
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}
