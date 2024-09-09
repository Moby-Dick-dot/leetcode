package main

func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}
	target := sum / 2

	dp := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], stones[i]+dp[j-stones[i]])
		}
	}

	return sum - dp[target] - dp[target]
}
