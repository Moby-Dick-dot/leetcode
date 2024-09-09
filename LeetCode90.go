package main

//import "sort"
//
//var (
//	path []int
//	used []bool
//	res  [][]int
//)
//
//func subsetsWithDup(nums []int) [][]int {
//	path = make([]int, 0)
//	used = make([]bool, len(nums))
//	res = make([][]int, 0)
//
//	if len(nums) == 0 {
//		return res
//	}
//
//	sort.Ints(nums)
//
//	backtracking(nums, 0, used)
//
//	return res
//}
//
//func backtracking(nums []int, startIndex int, used []bool) {
//	temp := make([]int, len(path))
//	copy(temp, path)
//	res = append(res, temp)
//
//	for i := startIndex; i < len(nums); i++ {
//		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
//			continue
//		}
//		path = append(path, nums[i])
//		used[i] = true
//		backtracking(nums, i+1, used)
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}
