//package main
//
//var (
//	path []int
//	res  [][]int
//)
//
//func findSubsequences(nums []int) [][]int {
//	path = make([]int, 0)
//	res = make([][]int, 0)
//
//	if len(nums) == 0 {
//		return res
//	}
//
//	backtracking(nums, 0)
//
//	return res
//}
//
//func backtracking(nums []int, startIndex int) {
//	if len(path) >= 2 {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//	}
//
//	mp := make(map[int]bool)
//
//	for i := startIndex; i < len(nums); i++ {
//		if (len(path) > 0 && nums[i] < path[len(path)-1]) || mp[nums[i]] {
//			continue
//		} //去重
//
//		path = append(path, nums[i])
//		mp[nums[i]] = true
//		backtracking(nums, i+1)
//		path = path[:len(path)-1]
//	}
//}
