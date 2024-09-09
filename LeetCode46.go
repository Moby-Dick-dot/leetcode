package main

//var (
//	path []int
//	res  [][]int
//	used []bool //标记那个元素已经使用过，然后跳过
//)
//
//func permute(nums []int) [][]int {
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
//
//}
//
//func backtracking(nums []int) {
//	if len(path) == len(nums) {
//		temp := make([]int, len(path))
//		copy(temp, path)
//		res = append(res, temp) //收获结果
//		return
//	}
//
//	for i := 0; i < len(nums); i++ { //注意：i从0开始，不是startindex
//		if used[i] == true {
//			continue
//		}
//
//		path = append(path, nums[i])
//		used[i] = true
//		backtracking(nums)
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}
