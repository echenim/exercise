package main

import "fmt"

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?

func main() {
	k := uniquePaths(3, 3)
	l := uniquePathrecursive(0, 1, 3, 3)
	fmt.Println(k, "  |   ", l)
}

//time complexity exponetial
//space complexity exponetial
//not a good solution
func uniquePathrecursive(i int, j int, n int, m int) int {
	if i == (n-1) && j == (m-1) {
		return 1
	}

	if i >= n || j >= m {
		return 0
	}

	return uniquePathrecursive(i+1, j, n, m) + uniquePathrecursive(i, j+1, n, m)

}

func uniquePaths(m int, n int) int {
	p := make([][]int, m)

	//fill the path of first row
	for i := 0; i < len(p); i++ {
		p[i] = make([]int, n)
		p[i][0] = 1
	}
	//fill the path ofr first column
	for i := 0; i < len(p[0]); i++ {
		p[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			p[i][j] = p[i-1][j] + p[i][j-1]
		}
	}

	return p[m-1][n-1]
}
