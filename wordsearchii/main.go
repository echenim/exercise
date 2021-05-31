package main

import "fmt"

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	//words := []string{"eat"}
	k := findWords(board, words)

	fmt.Println(k)
}

func findWords(board [][]byte, words []string) []string {
	var s []string

	for _, v := range words {
		//word := []byte(v)
		if help(board, v) {
			s = append(s, v)
		}
	}

	return s
}

func help(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0, board, word) {
				return true
			}
		}
	}
	return false
}

func backtrack(i int, j int, k int, board [][]byte, word string) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || word[k] != board[i][j] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	tmp := board[i][j]
	board[i][j] = 0
	found := backtrack(i+1, j, k+1, board, word) ||
		backtrack(i, j+1, k+1, board, word) ||
		backtrack(i-1, j, k+1, board, word) ||
		backtrack(i, j-1, k+1, board, word)

	board[i][j] = tmp
	return found

}
