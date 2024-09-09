package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	//初始化
	dp[0] = 1

	for i, l := 0, len(coins); i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]] //组合数必用公式
		}
	}

	return dp[amount]
}
