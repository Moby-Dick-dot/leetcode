//package main
//
//import "strings"
//
//var (
//	res [][]string
//)
//
//func solveNQueens(n int) [][]string {
//	res = make([][]string, 0)
//	chessboard := make([][]string, n) //初始化n行
//
//	for i := 0; i < n; i++ {
//		chessboard[i] = make([]string, n) //初始化n列
//	}
//
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			chessboard[i][j] = "."
//		}
//	}
//
//	if n == 0 {
//		return res
//	}
//
//	backtracking(chessboard, 0, n)
//
//	return res
//}
//
//func backtracking(chessboard [][]string, row int, n int) {
//	if row == n {
//		temp := make([]string, n)
//		for i, rowStr := range chessboard {
//			temp[i] = strings.Join(rowStr, "")
//		}
//		res = append(res, temp)
//		return
//	}
//
//	for i := 0; i < n; i++ {
//		if isValid(row, i, chessboard, n) {
//			chessboard[row][i] = "Q"
//			backtracking(chessboard, row+1, n)
//			chessboard[row][i] = "."
//		}
//	}
//}
//
//func isValid(row int, col int, chessboard [][]string, n int) bool {
//	for i := row; i >= 0; i-- {
//		if chessboard[i][col] == "Q" {
//			return false
//		}
//	}
//
//	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
//		if chessboard[i][j] == "Q" {
//			return false
//		}
//	}
//
//	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
//		if chessboard[i][j] == "Q" {
//			return false
//		}
//	}
//
//	return true
//}
