//package main
//
//var (
//	path []string
//	res  [][]string
//)
//
//func partition(s string) [][]string {
//	path = make([]string, 0)
//	res = make([][]string, 0)
//
//	if len(s) == 0 {
//		return make([][]string, 0)
//	}
//
//	backtracking(s, 0)
//	return res
//}
//
//func backtracking(s string, startindex int) {
//	if startindex == len(s) {
//		temp := make([]string, len(path))
//		copy(temp, path)
//		res = append(res, temp)
//		return
//	}
//
//	for i := startindex; i < len(s); i++ {
//		//判断是否为回文串，然后进行添加，不是就开启下一个循环
//		if isPalindrome(s, startindex, i) {
//			path = append(path, s[startindex:i+1])
//		} else {
//			continue
//		}
//		backtracking(s, i+1)
//		path = path[:len(path)-1]
//	}
//}
//
//func isPalindrome(s string, start int, end int) bool {
//	for i, j := start, end; i < j; i, j = i+1, j-1 {
//		if s[i] != s[j] {
//			return false
//		}
//	}
//	return true
//}
