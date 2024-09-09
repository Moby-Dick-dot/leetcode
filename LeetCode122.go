package main

//贪心
//func maxProfit(prices []int) int {
//	profit := 0
//
//	for i := 1; i < len(prices); i++ {
//		if prices[i]-prices[i-1] > 0 {
//			profit += prices[i] - prices[i-1]
//		}
//	}
//
//	return profit
//}

// 动态规划
//func maxProfit(prices []int) int {
//	dp := make([][]int, len(prices))
//
//	//dp[i][0] 持有股票手上最大的现金金额
//	//dp[i][1] 不持有股票手上的最大的现金金额
//	for i := 0; i < len(prices); i++ {
//		dp[i] = make([]int, 2)
//	}
//	dp[0][0] = -prices[0]
//	dp[0][1] = 0
//
//	for i := 1; i < len(prices); i++ {
//		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) //！！！和1不一样
//		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
//	}
//
//	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
//}
