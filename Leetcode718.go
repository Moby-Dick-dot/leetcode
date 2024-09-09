package main

func findLength(nums1 []int, nums2 []int) int {
	//dp[i][j] 表示以下标i-1为结尾的nums1，和以下标j-1为结尾的nums2，它们的最长重复子数组的长度
	dp := make([][]int, len(nums1)+1)

	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	res := dp[0][0]
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}

	return res
}
