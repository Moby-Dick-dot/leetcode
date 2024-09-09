//package main
//
//func solveSudoku(board [][]byte) {
//	backtracking(board)
//}
//
//func backtracking(board [][]byte) bool {
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[0]); j++ {
//			if board[i][j] == '.' {
//				for k := '0'; k < '9'; k++ {
//					if isValid(i, j, byte(k), board) {
//						board[i][j] = byte(k)
//						result := backtracking(board)
//						if result {
//							return true
//						}
//						board[i][j] = '.'
//					}
//				}
//				return false
//			}
//			return true
//		}
//	}
//}
//
//func isValid(row int, col int, k byte, board [][]byte) bool {
//	for i := 0; i < 9; i++ {
//		if board[i][col] == k {
//			return false
//		}
//	}
//
//	for i := 0; i < 9; i++ {
//		if board[row][i] == k {
//			return false
//		}
//	}
//
//	startrow := (row / 3) * 3
//	startcol := (col / 3) * 3
//
//	for i := startrow; i < startrow+3; i++ {
//		for j := startcol; j < startcol+3; j++ {
//			if board[i][j] == k {
//				return false
//			}
//		}
//	}
//
//	return true
//}
