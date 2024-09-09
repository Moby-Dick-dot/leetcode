//package main
//
//var (
//	path []int
//	res  [][]int
//	used []bool
//)
//
//func permuteUnique(nums []int) [][]int {
//	path = make([]int, 0)
//	res = make([][]int, 0)
//	used = make([]bool, len(nums))
//
//	if len(nums) == 0 {
//		return res
//	}
//
//	backtracking(nums)
//
//	return res
//}
//
//func backtracking(nums []int) {
//	if len(path) == len(nums) {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//		return
//	}
//
//	mp := make(map[int]bool)
//
//	for i := 0; i < len(nums); i++ {
//		if mp[nums[i]] || used[i] {
//			continue
//		}
//
//		path = append(path, nums[i])
//		mp[nums[i]] = true
//		used[i] = true
//		backtracking(nums)
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}
