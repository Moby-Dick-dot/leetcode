package main

func numDistinct(s string, t string) int {
	//dp[i][j] 以i-1为结尾的s中，有以j-1为结尾的t的个数
	//关键在递推公式      相等时，分为考虑和不考虑       不相等时，只能是不考虑，这一种情况
	//初始化要根据定义来

	l1 := len(s)
	l2 := len(t)

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i <= l1; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s[i-1] == t[j-1] {
				//如果相等
				//分为两种情况：考虑s[i-1] 和不考虑（即删除和不删除）
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j] //考虑  不考虑
			} else {
				//如果不相等 就不用考虑了
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[l1][l2]
}
