package main

func maxProfit(k int, prices []int) int {
	//总共有2K+1个状态 不操作 第一次持有 第一次不持有 。。。
	//dp[i][2k]

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	//初始化
	dp[0][0] = 0
	for j := 0; j < 2*k; j += 2 {
		dp[0][j+1] = -prices[0]
		dp[0][j+2] = 0
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}

	return dp[len(prices)-1][2*k]
}
