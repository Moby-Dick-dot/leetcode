package main

//func packages(m int, n int, space []int, value []int) int {
//	dp := make([][]int, m)
//
//	for i := 0; i < m; i++ {
//		dp[i] = make([]int, n+1)
//	}
//
//	//初始化
//	for i := 0; i < m; i++ {
//		dp[i][0] = 0
//	}
//
//	for j := 1; j < n+1; j++ {
//		if space[0] <= j {
//			dp[0][j] = value[0]
//		}
//	}
//
//	//开始遍历 先遍历物品，再遍历背包大小
//	for i := 1; i < m; i++ {
//		for j := 1; j < n+1; j++ {
//			if j-space[i] < 0 {
//				dp[i][j] = dp[i-1][j]
//			} else {
//				dp[i][j] = max(dp[i-1][j], value[i]+dp[i-1][j-space[i]])
//			}
//		}
//	}
//
//	return dp[m-1][n]
//}
//
//func max(i, j int) int {
//	if i <= j {
//		return j
//	} else {
//		return i
//	}
//}
//func main() {
//	var m, n int
//	fmt.Println("请输入材料的种类和行李箱的大小：")
//	fmt.Scan(&m, &n)
//
//	space := make([]int, m)
//	value := make([]int, m)
//
//	fmt.Println("请输入每种研究材料所占的空间：")
//	for i, _ := range space {
//		fmt.Scan(&space[i])
//	}
//
//	fmt.Println("请输入每种研究材料的价值：")
//	for i, _ := range value {
//		fmt.Scan(&value[i])
//	}
//
//	//fmt.Println("请输入材料的种类和行李箱的大小：", m, n)
//	//fmt.Println("请输入每种研究材料所占的空间：", space)
//	//fmt.Println("请输入每种研究材料的价值：", value)
//
//	fmt.Println(packages(m, n, space, value))
//}
