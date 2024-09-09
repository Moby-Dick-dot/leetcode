package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	//初始化
	for i := 0; i < len(dp); i++ {
		if obstacleGrid[i][0] == 1 {
			for j := i; j < len(dp); j++ {
				dp[j][0] = 0
			}
			break
		} else {
			dp[i][0] = 1
		}

	}

	for i := 0; i < len(dp[0]); i++ {

		if obstacleGrid[0][i] == 1 {
			for j := i; j < len(dp[0]); j++ {
				dp[0][j] = 0
			}
			break
		} else {
			dp[0][i] = 1
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
