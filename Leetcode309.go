package main

func maxProfit(prices []int) int {
	//dp[i][0] 第i天，持有股票时，手上的最大现金金额
	//dp[i][1] 第i天，不持有股票时，手上的最大现金金额

	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	//初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[1][0] = max(dp[0][0], -prices[1])
	dp[1][1] = max(dp[0][1], dp[0][0]+prices[1])

	for i := 2; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}
