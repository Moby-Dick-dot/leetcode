package main

import (
	"math"
	"sort"
)

//func combine(n int, k int) (res [][]int) {
//	path := make([]int, 0)
//	var backtracking func(n int, k int, startindex int)
//	backtracking = func(n int, k int, startindex int) {
//		if len(path) == k {
//			res = append(res, path) //将路径添加结果数组中
//			return
//		} //有返回，但不返回值
//
//		for i := startindex; i <= n; i++ { //横向遍历
//			path = append(path, i)
//			backtracking(n, k, i+1)
//			path = path[:len(path)-1]
//		}
//	}
//	backtracking(n, k, 1)
//	return res
//}

func runeReserve(runes []int) int {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	if len(runes) == 0 {
		return 0
	}

	if len(runes) == 1 {
		return 1
	}

	count := 1
	maxnum := math.MinInt

	for i := 1; i < len(runes); i++ {
		if runes[i]-runes[i-1] <= 1 {
			count++
		} else {
			maxnum = max(maxnum, count)
			count = 1
		}
	}

	return max(maxnum, count)
}
