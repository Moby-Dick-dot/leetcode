package main

func combinationSum3(k int, n int) (res [][]int) {
	path := make([]int, 0)
	var backtracking func(k int, n int, startindex int)
	backtracking = func(k int, n int, startindex int) {
		//剪枝
		if n < 0 {
			return
		}
		//终止条件
		if n == 0 && len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := startindex; i <= 9; i++ {
			path = append(path, i)
			backtracking(k, n-i, i+1) //使用减法来进行求和
			path = path[:len(path)-1]
		}
		return
	}
	backtracking(k, n, 1)
	return res
}
