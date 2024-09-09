package main

func minDistance(word1 string, word2 string) int {
	//dp[i][j] 将下标0-i-1的 word1 转换成 下标0-j-1的word2 所使用的最小操作数 是dp[i][j]
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	//初始化
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1) //删除(和插入本质上是一样的) 替换
			}
		}
	}

	return dp[l1][l2]
}
