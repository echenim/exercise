package main

// Given a robot cleaner in a room modeled as a grid.
// Each cell in the grid can be empty or blocked.
// The robot cleaner with 4 given APIs can move forward,
// turn left or turn right. Each turn it made is 90 degrees.
// When it tries to move into a blocked cell, its bumper sensor
// detects the obstacle and it stays on the current cell.

// Design an algorithm to clean the entire room using only the 4 given APIs shown below.

// interface Robot {
//   // returns true if next cell is open and robot moves into the cell.
//   // returns false if next cell is obstacle and robot stays on the current cell.
//   boolean move();

//   // Robot will stay on the same cell after calling turnLeft/turnRight.
//   // Each turn will be 90 degrees.
//   void turnLeft();
//   void turnRight();

//   // Clean the current cell.
//   void clean();
// }

// Input:
// room = [
//   [1,1,1,1,1,0,1,1],
//   [1,1,1,1,1,0,1,1],
//   [1,0,1,1,1,1,1,1],
//   [0,0,0,1,0,0,0,0],
//   [1,1,1,1,1,1,1,1]
// ],
// row = 1,
// col = 3

// Explanation:
// All grids in the room are marked by either 0 or 1.
// 0 means the cell is blocked, while 1 means the cell is accessible.
// The robot initially starts at the position of row=1, col=3.
// From the top left corner, its position is one row below and three columns right.

func main() {
	rm := [][]int{
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1}}
}

type Robot struct{}

type location struct {
	x int
	y int
}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func track(robot *Robot, r int, c int, d int, visited map[location]bool) {
	currentLocation := location{r, c}
	visited[currentLocation] = true
	robot.Clean()

	for i := 0; i < 4; i++ {
		dx := r + directions[d][0]
		dy := c + directions[d][1]

		if !visited[location{dx, dy}] && robot.Move() {
			track(robot, dx, dy, d, visited)
		}

		d = (d + 1) % 4
		robot.TurnRight()
	}
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnLeft()
	robot.TurnLeft()
}

func cleanRoom(robot *Robot) {
	visited := make(map[location]bool)
	track(robot, 0, 0, 0, visited)
}

func (robot *Robot) Move() bool { return true }
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}
func (robot *Robot) Clean()     {}
