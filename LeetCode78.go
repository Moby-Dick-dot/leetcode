//package main
//
//var (
//	path []int
//	res  [][]int
//)
//
//func subsets(nums []int) [][]int {
//	path = make([]int, 0)
//	res = make([][]int, 0)
//
//	backtracking(nums, 0)
//	return res
//}
//
//func backtracking(nums []int, startindex int) {
//	tmp := make([]int, len(path))
//	copy(tmp, path)
//	res = append(res, tmp) // 将当前路径添加到结果中
//
//	for i := startindex; i < len(nums); i++ {
//		path = append(path, nums[i]) // 做选择
//		backtracking(nums, i+1)      // 递归
//		path = path[:len(path)-1]    // 撤销选择
//	}
//}
