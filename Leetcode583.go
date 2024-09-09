package main

func minDistance(word1 string, word2 string) int {
	//dp[i][j] 下标0-i-1的word1 和 下标0-j-1的word2 相同最小的操作步数
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	//根据定义初始化
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] //不用删除操作，所以操作次数和dp[i-1][j-1]相同
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+2) //分为3种情况，第一种删除word[i-1],第二种删除word[j-1],第三种既删除word[i-1]，又删除word[j-1]
			}
		}
	}

	return dp[l1][l2]

}
