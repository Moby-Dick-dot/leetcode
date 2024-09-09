//package main
//
//import (
//	"strconv"
//	"strings"
//)
//
//var (
//	path []string
//	res  []string
//)
//
//func restoreIpAddresses(s string) []string {
//	path = make([]string, 0)
//	res = make([]string, 0)
//
//	if len(s) == 0 {
//		return res
//	}
//
//	backtracking(s, 0)
//	return res
//}
//
//func backtracking(s string, startindex int) {
//	if len(path) == 4 {
//		if startindex == len(s) {
//			str := strings.Join(path, ".")
//			res = append(res, str)
//		}
//		return
//	}
//
//	for i := startindex; i < len(s); i++ {
//		if isValid(s[startindex:i+1], startindex, i) {
//			path = append(path, s[startindex:i+1])
//			backtracking(s, i+1)
//			path = path[:len(path)-1]
//		} else {
//			break
//		}
//	}
//}
//
//func isValid(s string, start int, end int) bool {
//	if s[0] == '0' && start != end {
//		return false
//	}
//	num, _ := strconv.Atoi(s)
//	if num > 255 || num < 0 {
//		return false
//	}
//	return true
//}
