package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)

	for a := 0; a < m+1; a++ {
		dp[a] = make([]int, n+1)
	}

	for k := 0; k < len(strs); k++ {
		zeronum := 0
		onenum := 0
		for _, c := range strs[k] {
			if c == '0' {
				zeronum += 1
			} else {
				onenum += 1
			}
		}

		for i := m; i >= zeronum; i-- {
			for j := n; j >= onenum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeronum][j-onenum]+1)
			}
		}
	}

	return dp[m][n]
}
