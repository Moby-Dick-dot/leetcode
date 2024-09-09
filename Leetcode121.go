package main

//func maxProfit(prices []int) int {
//	//方法一：暴力，两层循环找出当天买入然后之后卖出的最大利润，最后再求个最大值
//	//方法二：dp
//	//每一天都有两种状态，持有股票和不持有股票（买入股票和卖出股票包含不了所有情况）
//	//dp[i][0] 表示持有股票时，手中的最大现金（利润）金额
//	//dp[i][1] 表示不持有股票时，手中的最大现金金额
//	//状态转移方程 dp[i][0] = max(dp[i-1][0],-price[i])  dp[i][1] = max(dp[i-1][1],dp[i][0]+price[i])
//	//初始化 i需要从i-1得到 dp[0][0] = -price[0]   dp[0][1] = 0
//
//	dp := make([][]int, len(prices))
//
//	for i := 0; i < len(prices); i++ {
//		dp[i] = make([]int, 2)
//	}
//	dp[0][0] = -prices[0]
//	dp[0][1] = 0
//
//	for i := 1; i < len(prices); i++ {
//		dp[i][0] = max(dp[i-1][0], -prices[i])           //持有股票的状态，可以由前一天持有转过来，也可以由前一天不持有转过来
//		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) //不持有股票的状态，可以由前一天不持有股票的状	态转过来，也可以由前一天持有股票的状态转过来
//	}
//
//	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
//}
