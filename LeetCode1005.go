//package main
//
//import "sort"
//
//func sortByAbsoluteValue(arr []int) {
//	sort.Slice(arr, func(i, j int) bool {
//		return abs(arr[i]) > abs(arr[j])
//	})
//}
//
//func abs(num int) int {
//	if num < 0 {
//		return -num
//	} else {
//		return num
//	}
//}
//
//func largestSumAfterKNegations(nums []int, k int) int {
//	//第一步 将数组按照绝对值从大到小进行排序
//	sortByAbsoluteValue(nums)
//
//	//第二步：从前往后对负数取反
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < 0 && k > 0 {
//			nums[i] *= -1
//			k--
//		}
//	}
//
//	//第三步：如果k不为0，那么重复对一个数取反
//	if k%2 == 1 {
//		nums[len(nums)-1] *= -1
//	}
//
//	//第四步：进行求和
//	result := 0
//	for i := 0; i < len(nums); i++ {
//		result += nums[i]
//	}
//
//	return result
//}
