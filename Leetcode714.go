package main

func maxProfit(prices []int, fee int) int {
	//dp[i][0] 第i天，持有股票时，手上的最大现金金额
	//dp[i][1] 第i天，不持有股票时，手上的最大现金金额

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	//初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[len(prices)-1][1]
}
