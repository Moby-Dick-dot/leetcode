package main

import "math"

func coinChange(coins []int, amount int) int {

	//硬币无限，属于多重背包，要从小到大开始遍历
	//dp[j]含义 表示装满容量为j的背包，所能存放的最小价值(和之前的不一样，之前是容量为j的背包，所能存放的最大价值，没有装满的意思)
	//dp[j]含义凑足总额为j所需钱币的最少个数为dp[j]

	//dp数组如何初始化
	//首先凑足总金额为0所需钱币的个数一定是0，那么dp[0] = 0;
	//
	//其他下标对应的数值呢？
	//
	//考虑到递推公式的特性，dp[j]必须初始化为一个最大的数，否则就会在min(dp[j - coins[i]] + 1, dp[j])比较的过程中被初始值覆盖。
	//
	//所以下标非0的元素都是应该是最大值。
	dp := make([]int, amount+1)

	//初始化 当硬币金额和背包容量相等时，刚好放下一个，价值为1
	dp[0] = 0
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	for i, l := 0, len(coins); i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
