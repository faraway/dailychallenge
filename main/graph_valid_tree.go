package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
You have a graph of n nodes labeled from 0 to n - 1.
You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates
that there is an undirected edge between nodes ai and bi in the graph.

Return true if the edges of the given graph make up a valid tree, and false otherwise.



Example 1:
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
             0
           / | \
          1  2  3
          |
          4
Output: true


Example 2:
Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
          0 -- 1 -- 2
               | \  |
               |  \ |
               4   3
Output: false

Constraints:
1 <= 2000 <= n
0 <= edges.length <= 5000
edges[i].length == 2
0 <= ai, bi < n
ai != bi
There are no self-loops or repeated edges.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", validTree(5, [][]int{{0,1}, {0,2}, {0,3}, {1,4}})) //True
	fmt.Println("Answer is:", validTree(5, [][]int{{0,1}, {1,2}, {2,3}, {1,3}, {1,4}})) //False
}

func validTree(n int, edges [][]int) bool {
	ds := NewDisjointSet(n)
	for _, edge := range edges {
		if ds.Connected(edge[0], edge[1]) {
			//these two nodes are already (indirectly) connected, connect them here will cause a loop, thus not a tree
			return false
		}
		ds.Union(edge[0], edge[1])
	}
	//now we make sure all nodes are connected
	root := ds.Find(ds.Parent[0])
	for _, node := range ds.Parent {
		if ds.Find(node) != root {
			return false
		}
	}
	return true
}