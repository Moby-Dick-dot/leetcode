package main

//import "sort"
//
//func merge(intervals [][]int) [][]int {
//	//按照左边界对区间进行排序
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][0] < intervals[j][0]
//	})
//
//	if len(intervals) == 0 {
//		return [][]int{}
//	}
//	//不断更新左,右边界
//
//	result := make([][]int, 0)
//
//	for i := 1; i < len(intervals); i++ {
//		if intervals[i][0] <= intervals[i-1][1] {
//			intervals[i][0] = intervals[i-1][0]
//			intervals[i][1] = max(intervals[i][1], intervals[i-1][1])
//		} else {
//			result = append(result, intervals[i-1])
//		}
//	}
//
//	result = append(result, intervals[len(intervals)-1])
//
//	return result
//}
