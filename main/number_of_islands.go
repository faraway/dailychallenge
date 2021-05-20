package main

import "fmt"

/**
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.


*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", numIslands([][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	})) //3
}

func numIslands(grid [][]byte) int {
	numOfIslands := 0

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] == byte('1') {
				numOfIslands++
				exploreIsland(&grid,i,j)
			}
		}
	}

	return numOfIslands
}

func exploreIsland(grid *[][]byte, i int, j int) {
	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) {
		return
	}
	if (*grid)[i][j] == byte('A') || (*grid)[i][j] == byte('0') { //visited nodes or water
		return
	}
	(*grid)[i][j] = byte('A')
	exploreIsland(grid, i-1, j)
	exploreIsland(grid, i+1, j)
	exploreIsland(grid, i, j-1)
	exploreIsland(grid, i, j+1)
}
