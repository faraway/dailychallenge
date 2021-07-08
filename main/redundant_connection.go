package main

import "fmt"

/**
https://leetcode.com/problems/redundant-connection/

In this problem, a tree is an undirected graph that is connected and has no cycles.

You are given a graph that started as a tree with n nodes labeled from 1 to n, with one additional edge added.
The added edge has two different vertices chosen from 1 to n, and was not an edge that already existed.
The graph is represented as an array edges of length n where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the graph.

Return an edge that can be removed so that the resulting graph is a tree of n nodes.
If there are multiple answers, return the answer that occurs last in the input.



Example 1:
  1 -----  2
  |       /
  |      /
  |     /
  |    /
  |   /
  |  /
   3
Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]


Example 2:
  2 ---- 1 ---- 5
  |      |
  |      |
  |      |
  3 ---- 4

Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]


Constraints:
n == edges.length
3 <= n <= 1000
edges[i].length == 2
1 <= ai < bi <= edges.length
ai != bi
There are no repeated edges.
The given graph is connected.

*/

func main() {
	fmt.Println("initializing test data...")

	fmt.Println("Answer is:", findRedundantConnection([][]int{{1,2}, {2,3}, {3,4}, {1,4}, {1,5}})) //[1,4]
}

/**
 TODO:
 Alternative solution : https://leetcode.com/problems/redundant-connection/solution/ DFS
 Incrementally build up the graph while for each new edge seen, check if it makes a loop. This way "find the last edge in the list" is a given.
 */

/**
 The following method tries to build up the entire graph first, and then detect the loop, and then figure out "the last edge in the list"
 */
func findRedundantConnection(edges [][]int) []int {
	edgesIndexMap := make(map[[2]int]int)
	graph := make(map[int][]int)

	for idx, edge := range edges {
		//each edge A--B implies A->B and B->A
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		edgesIndexMap[[2]int{edge[0], edge[1]}] = idx
	}

	visited := make(map[int]bool)
	//use edges[0][0] because it's all connected, doesn't really matter which 'node' we start with
	loopPath := redudantConnectionDfs(graph, &visited, edges[0][0], []int{})

	maxIdx := 0
	var result [2]int
	//loop path will be something look like this [5,1,2,3,4,1]
	for i:= len(loopPath)-2; i>=0; i-- {
		e1 := [2]int {loopPath[i], loopPath[i+1]}
		e2 := [2]int {loopPath[i+1], loopPath[i]}
		candidate := e1
		idx, existing := edgesIndexMap[e1]
		if !existing {
			idx = edgesIndexMap[e2]
			candidate = e2
		}
		if idx > maxIdx {
			maxIdx = idx
			result = candidate
		}
		if loopPath[i] == loopPath[len(loopPath)-1] { //do not go back further as this is the beginning of the loop
			break
		}
	}
	return result[:]

}

func redudantConnectionDfs(graph map[int][]int, visited *map[int]bool, node int, path []int) []int {
	if (*visited)[node] {
		return append(path, node)
	} else {
		(*visited)[node] = true
		for _, adjacentNode := range graph[node] {
			if len(path) > 0 && adjacentNode == path[len(path)-1] { //do not go back to the previous node
				continue
			}
			loopFound := redudantConnectionDfs(graph, visited, adjacentNode, append(path, node))
			if loopFound != nil {
				return loopFound
			}
		}
		return nil
	}
}