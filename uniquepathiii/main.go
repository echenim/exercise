package main

// On a 2-dimensional grid, there are 4 types of squares:

// 1 represents the starting square.  There is exactly one starting square.
// 2 represents the ending square.  There is exactly one ending square.
// 0 represents empty squares we can walk over.
// -1 represents obstacles that we cannot walk over.
// Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.

// Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
// Output: 2

func main() {

}

func uniquePathsIII(grid [][]int) int {
	x, y, n := -1, -1, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				n++
			}
			if grid[i][j] == 1 {
				x, y, n = i, j, n+1
			}
		}
	}
	return dfs(&grid, x, y, n)
}

func dfs(grid *[][]int, x, y, n int) int {
	if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[0]) || (*grid)[x][y] == -1 {
		return 0
	}
	if (*grid)[x][y] == 2 {
		if n == 0 {
			return 1
		}
		return 0
	}
	(*grid)[x][y] = -1
	paths := dfs(grid, x+1, y, n-1) +
		dfs(grid, x-1, y, n-1) +
		dfs(grid, x, y+1, n-1) +
		dfs(grid, x, y-1, n-1)
	(*grid)[x][y] = 0
	return paths
}
