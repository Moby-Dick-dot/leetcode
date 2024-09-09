package main

var (
	path []int
	res  [][]int
)

func combinationSum(candidates []int, target int) [][]int {
	path = make([]int, 0)
	res = make([][]int, 0)

	if len(candidates) == 0 {
		return res
	}

	dfs(candidates, target, 0)

	return res
}

func dfs(candidates []int, target int, startindex int) {
	if target < 0 {
		return
	}

	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}

	for i := startindex; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfs(candidates, target-candidates[i], i) // 关键点:不用i+1了，表示可以重复读取当前的数
		path = path[:len(path)-1]
	}

	return
}
