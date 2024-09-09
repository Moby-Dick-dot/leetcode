//package main
//
//import "sort"
//
//func findContentChildren(g []int, s []int) int {
//	count := 0
//
//	sort.Ints(g)
//	sort.Ints(s)
//
//	for i, j := 0, 0; i < len(g) && j < len(s); {
//		if s[j] < g[i] {
//			j++
//		} else {
//			i++
//			j++
//			count++
//		}
//	}
//
//	return count
//}
