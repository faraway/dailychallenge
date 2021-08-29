package main

import (
	. "dailychallenge/utils"
	"fmt"
)

/**
https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/

You have a graph of n nodes. You are given an integer n and an array edges where edges[i] = [ai, bi]
indicates that there is an edge between ai and bi in the graph.

Return the number of connected components in the graph.

Example 1:
Input: n = 5, edges = [[0,1],[1,2],[3,4]]
    0 -- 1   3
         |   |
         2   4

Output: 2


Example 2:
Input: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
    0 -- 1   3
         | / |
         2   4

Output: 1


Constraints:

1 <= n <= 2000
1 <= edges.length <= 5000
edges[i].length == 2
0 <= ai <= bi < n
ai != bi
There are no repeated edges.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", countComponents(5, [][]int{{0,1}, {1,2}, {3,4}})) //2
	fmt.Println("Answer is:", countComponents(5, [][]int{{0,1}, {1,2}, {2,3}, {3,4}})) //1
}

func countComponents(n int, edges [][]int) int {
	ds := NewDisjointSet(n)
	for _, edge := range edges {
		ds.Union(edge[0], edge[1])
	}
	set := make(map[int]int)
	for node := 0; node < n; node++ {
		set[ds.Find(node)]++
	}
	return len(set)
}


/**
	We can also use DFS or BFS
	1. turn edges into adjacent list
	2. for node 1..n (that is not visited yet)
			dfs (node) and mark them visited (this traverses all the nodes connected to this node, thus one 'component')
			count ++
    3. return count
 */
