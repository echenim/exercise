package main

import "fmt"

func main() {
	sa := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}, {0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	k := uniquePathsWithObstacles(sa)
	fmt.Println(k)

}

func uniquePathsWithObstacles(p [][]int) int {
	//is if starting point is an obstacle
	if p[0][0] == 1 {
		return 0
	}

	p[0][0] = 1

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[0]); j++ {
			if p[i][j] == 1 && (i+j) > 0 {
				p[i][j] = -1
			}
			if j > 0 && p[0][j-1] > 0 && p[0][j] != -1 {
				p[0][j] = 1
			}
		}
		if i > 0 && p[i-1][0] > 0 && p[i][0] != -1 {
			p[i][0] = 1
		}
	}

	for i := 1; i < len(p); i++ {
		for j := 1; j < len(p[0]); j++ {
			if p[i][j] == -1 {
				continue
			}
			if p[i-1][j] == -1 && p[i][j-1] == -1 {
				continue
			}
			if p[i-1][j] != -1 && p[i][j-1] == -1 {
				p[i][j] = p[i-1][j]
			} else if p[i-1][j] == -1 && p[i][j-1] != -1 {
				p[i][j] = p[i][j-1]
			} else {
				p[i][j] = p[i-1][j] + p[i][j-1]
			}
		}
	}
	r := p[len(p)-1][len(p[0])-1]
	if r < 0 {
		return 0
	}
	return r
}
