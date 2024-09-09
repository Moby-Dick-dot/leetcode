package main

import "fmt"

func integerBreak(n int) int {
	//dp[i]表示拆分数字i，得到的最大乘积是dp[i]
	dp := make([]int, n+1)

	//初始化
	dp[0] = 0 //无意义
	dp[1] = 0 //无意义
	dp[2] = 1

	//状态转移
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(j*(i-j), j*dp[i-j], dp[i])
		}
	}

	fmt.Println(dp)
	return dp[n]

}

func main() {
	integerBreak(6)
}
