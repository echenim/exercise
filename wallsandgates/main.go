package main

import "fmt"

func main() {
	sa := [][]int{{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647}}
	wallsAndGates(sa)
	fmt.Println(sa)
}

func wallsAndGates(rooms [][]int) {

	for i := 0; i < len(rooms); i++ {

		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == 0 {
				dsf(i, j, 0, rooms)
			}
		}
	}
}

func dsf(i int, j int, distance int, rooms [][]int) {
	//check of we are still within the grid else stop traversing
	if i < 0 || i >= len(rooms) || j < 0 || j >= len(rooms[i]) || rooms[i][j] < distance {
		return
	}

	rooms[i][j] = distance

	dsf(i+1, j, distance+1, rooms)
	dsf(i-1, j, distance+1, rooms)
	dsf(i, j+1, distance+1, rooms)
	dsf(i, j-1, distance+1, rooms)
}
