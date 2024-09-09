package main

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	//dp[i][j] 长度为[0,i-1]的nums1，与长度为[0,j-1]的nums2的最大连线数是dp[i][j]
	l1 := len(nums1)
	l2 := len(nums2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[l1][l2]
}
