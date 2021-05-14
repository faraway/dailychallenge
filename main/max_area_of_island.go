package main

import "fmt"

/**
https://leetcode.com/problems/max-area-of-island/
You are given an m x n binary matrix grid. An island is a group of 1's (representing land)
connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.

Example 1:
Input: grid =
[
	[0,0,1,0,0,0,0,1,0,0,0,0,0],
	[0,0,0,0,0,0,0,1,1,1,0,0,0],
	[0,1,1,0,1,0,0,0,0,0,0,0,0],
	[0,1,0,0,1,1,0,0,1,0,1,0,0],
	[0,1,0,0,1,1,0,0,1,1,1,0,0],
	[0,0,0,0,0,0,0,0,0,0,1,0,0],
	[0,0,0,0,0,0,0,1,1,1,0,0,0],
	[0,0,0,0,0,0,0,1,1,0,0,0,0]
]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.

Example 2:
Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] is either 0 or 1.
*/

func main() {
	fmt.Println("initializing test data...")

	islands1 := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println("Answer is:", maxAreaOfIsland(islands1)) //6
	fmt.Println("Answer is:", maxAreaOfIsland([][]int{{0,0,0,0,0,0,0,0}})) //0

}

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxArea := 0
	for i := 0; i < m; i++ {
		for j:=0; j < n; j++ {
			if grid[i][j] > 0 {
				area := backtrackIsland(i,j,&grid)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func backtrackIsland(i int, j int, grid *[][]int) int {
	m := len(*grid)
	n := len((*grid)[0])
	//check bounds
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	if (*grid)[i][j] <= 0 {
		return 0
	}
	//mark it to -1 so that we know we have traversed this spot
	(*grid)[i][j] = -1
	return 1 + backtrackIsland(i-1, j, grid) + backtrackIsland(i+1, j, grid) +
		backtrackIsland(i, j-1, grid) + backtrackIsland(i, j+1, grid)
}