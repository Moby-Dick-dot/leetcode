package main

//import "sort"
//
//func eraseOverlapIntervals(intervals [][]int) int {
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][0] < intervals[j][0]
//	})
//
//	if len(intervals) == 0 {
//		return 0
//	}
//
//	count := 0
//
//	for i := 1; i < len(intervals); i++ {
//		if intervals[i][0] >= intervals[i-1][1] {
//
//		} else {
//			count++
//			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
//		}
//	}
//
//	return count
//}
