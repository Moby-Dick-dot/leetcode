package main

//func canPartition(nums []int) bool {
//	numslen := len(nums)
//	sum := 0
//	for i := 0; i < numslen; i++ {
//		sum += nums[i]
//	}
//
//	if sum % 2 ==1{
//		return false
//	}
//
//	target := sum / 2 //背包大小
//	dp := make([]int, target+1)
//
//	//初始化全部为0
//
//	//遍历
//	for i := 0; i < numslen; i++ {
//		for j := target; j >= nums[i]; j-- {
//			dp[j] = max(dp[j], nums[i]+dp[j-nums[i]])
//		}
//	}
//
//	fmt.Println(dp)
//
//	return dp[len(dp)-1] == target
//}
//
//func main() {
//	nums := []int{1, 5, 11, 5}
//
//	canPartition(nums)
//}
