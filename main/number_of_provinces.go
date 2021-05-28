package main

import "fmt"

/**
There are n cities. Some of them are connected, while some are not.
If city a is connected directly with city b, and city b is connected directly with city c,
then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city
and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.

Example 1:
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

Example 2:
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3

Constraints:

1 <= n <= 200
n == isConnected.length
n == isConnected[i].length
isConnected[i][j] is 1 or 0.
isConnected[i][i] == 1
isConnected[i][j] == isConnected[j][i]
 */

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", findCircleNum([][]int {{1,1,0},{1,1,0},{0,0,1}})) //2
	fmt.Println("Answer is:", findCircleNum([][]int {{1,0,0},{0,1,0},{0,0,1}})) //3
	fmt.Println("Answer is:", findCircleNum([][]int {{1,0,0,1},{0,1,1,0},{0,1,1,1},{1,0,1,1}})) //1
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	result := 0
	visited := make([]bool, n)

	for i := 0; i < n; i ++ {
		if !visited[i] {
			result ++
			//checker[i] = true
			exploreConnection(isConnected, &visited, i)
		}
	}

	return result
}

/**
 DFS. BSF should also work.
 */
func exploreConnection(isConnected [][]int, visited *[]bool, i int) {
	if (*visited)[i] {
		return
	}
	(*visited)[i] = true

	for j := 0; j < len(*visited); j++ {
		if i!=j && isConnected[i][j] == 1 {
			exploreConnection(isConnected, visited, j)
		}
	}
}