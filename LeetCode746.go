package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	//dp[i]表示爬到第i层阶梯，所需要的最小花费
	dp := make([]int, len(cost)+1)

	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	minCostClimbingStairs(cost)
}
