//package main
//
//var (
//	mp   []string
//	path []byte
//	res  []string
//)
//
//func letterCombinations(digits string) []string {
//	mp = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
//	path = make([]byte, 0)
//	res = make([]string, 0)
//
//	if len(digits) == 0 {
//		return res
//	}
//
//	dfs(digits, 0)
//
//	return res
//}
//
//func dfs(digits string, index int) {
//	if len(path) == len(digits) {
//		temp := string(path)
//		res = append(res, temp)
//		return
//	}
//
//	digit := digits[index] - '0'
//	letter := mp[digit]
//
//	for i := 0; i < len(letter); i++ {
//		path = append(path, letter[i])
//		dfs(digits, index+1)
//		path = path[:len(path)-1]
//	}
//	return
//}
