package main

func isSubsequence(s string, t string) bool {
	//dp[i][j] 以下标i-1结尾的字符串s，和以下标j-1结尾的字符串t，相同子序列的长度为dp[i][j]
	l1 := len(s)
	l2 := len(t)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[l1][l2] == l1
}
