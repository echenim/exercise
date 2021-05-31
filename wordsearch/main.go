package main

import "fmt"

func main() {
	b := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	k := isexist(b, word)
	fmt.Println(k)
}

func exist(board [][]byte, word string) bool {
	nrow := len(board)
	if nrow == 0 {
		return false
	}
	ncol := len(board[0])
	words := []byte(word)

	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			if backtrack(i, j, 0, board, words) {
				return true
			}
		}
	}

	return false
}

func backtrack(i int, j int, count int, board [][]byte, word []byte) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] == word[count] {
		if count == len(word)-1 {
			return true
		}

		tmp := board[i][j]
		board[i][j] = 0
		found := backtrack(i+1, j, count+1, board, word) ||
			backtrack(i, j+1, count+1, board, word) ||
			backtrack(i-1, j, count+1, board, word) ||
			backtrack(i, j-1, count+1, board, word)

		board[i][j] = tmp
		return found
	}

	return false
}

func isexist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, i, j, 0, word) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, i int, j int, count int, word string) bool {
	if count == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board) || board[i][j] != word[count] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = 0
	found := dfs(board, i+1, j, count+1, word) || dfs(board, i-1, j, count+1, word) ||
		dfs(board, i, j+1, count+1, word) || dfs(board, i, j-1, count+1, word)
	board[i][j] = tmp

	return found
}
